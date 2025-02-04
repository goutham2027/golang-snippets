package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	type s struct {
		Id string `json:"id,omitempty"`
	}
	j := []byte(`{"id": null}`)
	s1 := &s{}
	err := json.Unmarshal(j, s1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("%+v", s1)
}
