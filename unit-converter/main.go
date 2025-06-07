package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

type ConversionRequest struct {
	Type  string  `json:"type"`
	Value float64 `json:"value"`
	From  string  `json:"from"`
	To    string  `json:"to"`
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))
	err := tmpl.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func ConvertHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if ct := r.Header.Get("Content-Type"); ct != "application/json" {
		http.Error(w, "Expected JSON", http.StatusUnsupportedMediaType)
		return
	}

	var req ConversionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	response := map[string]interface{}{
		"status":  "success",
		"message": "Succesful convertion",
		"value":   FloatToString(convert(req)),
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", HomeHandler)

	http.HandleFunc("/api/convert", ConvertHandler)

	http.ListenAndServe(":8080", nil)
}

func FloatToString(f float64) string {
	s := fmt.Sprintf("%.10f", f)
	s = strings.TrimRight(s, "0")
	s = strings.TrimRight(s, ".")
	return s
}

func convert(cr ConversionRequest) float64 {
	switch cr.Type {
	case "length":
		return convertLength(cr)
	case "weight":
		return convertWeight(cr)
	case "temperature":
		return convertTemperature(cr)
	default:
		return 0
	}
}

func convertLength(cr ConversionRequest) float64 {
	var answer float64

	switch cr.From {
	case "millimeter":
		answer = cr.Value * 0.001
	case "centimeter":
		answer = cr.Value * 0.01
	case "meter":
		answer = cr.Value
	case "kilometer":
		answer = cr.Value * 1000
	case "inch":
		answer = cr.Value * 0.0254
	case "foot":
		answer = cr.Value * 0.3048
	case "yard":
		answer = cr.Value * 0.9144
	case "mile":
		answer = cr.Value * 1609.34
	default:
		answer = cr.Value
	}

	switch cr.To {
	case "millimeter":
		answer = answer / 0.001
	case "centimeter":
		answer = answer / 0.01
	case "meter":
	case "kilometer":
		answer = answer / 1000
	case "inch":
		answer = answer / 0.0254
	case "foot":
		answer = answer / 0.3048
	case "yard":
		answer = answer / 0.9144
	case "mile":
		answer = answer / 1609.34
	default:
	}

	return answer
}

func convertWeight(cr ConversionRequest) float64 {
	var answer float64

	switch cr.From {
	case "milligram":
		answer = cr.Value * 0.001
	case "gram":
		answer = cr.Value
	case "kilogram":
		answer = cr.Value * 1000
	case "ounce":
		answer = cr.Value * 28.3495
	case "pound":
		answer = cr.Value * 453.592
	default:
		answer = cr.Value
	}

	switch cr.To {
	case "milligram":
		answer = cr.Value / 0.001
	case "gram":
	case "kilogram":
		answer = cr.Value / 1000
	case "ounce":
		answer = cr.Value / 28.3495
	case "pound":
		answer = cr.Value / 453.592
	default:
	}

	return answer
}

func convertTemperature(cr ConversionRequest) float64 {
	var answer float64

	switch cr.From {
	case "celsius":
		answer = cr.Value
	case "fahrenheit":
		answer = (cr.Value - 32) * 5 / 9
	case "kelvin":
		answer = cr.Value - 273.15
	default:
		answer = cr.Value
	}

	switch cr.To {
	case "celsius":
	case "fahrenheit":
		answer = (cr.Value * 9 / 5) + 32
	case "kelvin":
		answer = cr.Value + 273.15
	default:
	}

	return answer
}
