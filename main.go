package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Person struct {
	Name       string `json:"customerName"`
	Age        int    `json:"age,omitempty"`
	CreditCard string `json:"-"`
}

func main() {
	p := Person{Name: "Tom", CreditCard: "super secret"}
	pBytes, err := json.Marshal(p)
	log.Print(err)
	log.Print(string(pBytes))

	p2AsRawJson, err3 := ioutil.ReadFile("john.json")
	log.Print(err3)

	var p2 Person
	err2 := json.Unmarshal(p2AsRawJson, &p2)
	log.Print(err2)
	log.Printf("%+v", p2)
}
