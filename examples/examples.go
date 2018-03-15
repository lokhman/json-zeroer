package main

import (
	"encoding/json"
	"time"

	_ "github.com/lokhman/json-zeroer"
)

type ZeroObject struct {
	A string
}

func (_ ZeroObject) IsZero() bool {
	return true
}

type Object struct {
	A string     `json:"a,omitempty"`
	B time.Time  `json:"b,omitempty"`
	C time.Time  `json:"c,omitempty"`
	D ZeroObject `json:"d,omitempty"`
}

func main() {
	o := Object{A: "Test", B: time.Now(), D: ZeroObject{"Test"}}
	data, err := json.Marshal(o)
	if err != nil {
		panic(err)
	}
	println(string(data))
}
