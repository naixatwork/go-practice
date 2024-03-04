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
	resp, err := http.Get("https://api.github.com/users/naixatwork")

	if err != nil {
		log.Fatalf("error: %s", err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("error: %s", resp.Status)
	}

	fmt.Printf("Content-Type: %s\n", resp.Header.Get("Content-Type"))
	if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
		log.Fatalf("error: can't copy - %s", err)
	}

	var r Reply
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&r); err != nil {
		log.Fatalf("error: can't decode - %s", err)
	}

	fmt.Printf("%#v\n", r)
}

type Reply struct {
	Name         string
	Public_Repos int `json:"public_repos"`
}
