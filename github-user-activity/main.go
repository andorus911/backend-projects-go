package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {

	username := os.Args[1]

	if len(username) == 0 {
		err := fmt.Errorf("no arguments")
		log.Fatalln(err)
		return
	}

	resp, err := http.Get("https://api.github.com/users/" + username + "/events")
	if err != nil {
		log.Fatalln(err)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
		return
	}

	var objmap []map[string]interface{}
	if err := json.Unmarshal(body, &objmap); err != nil {
		log.Fatal(err)
	}
	for _, v := range objmap {
		fmt.Println(v["id"])
		// switch
	}
}
