# json-zeroer - Zero-friendly JSON patch for Go

[![Build Status](https://travis-ci.org/lokhman/json-zeroer.svg?branch=master)](https://travis-ci.org/lokhman/json-zeroer)
[![codecov](https://codecov.io/gh/lokhman/json-zeroer/branch/master/graph/badge.svg)](https://codecov.io/gh/lokhman/json-zeroer)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

Implementation of Go issue [#11939](https://github.com/golang/go/issues/11939) that uses memory hack known as
"monkey patching" presented by [Bouke van der Bijl](https://github.com/bouk/monkey).

## Install

	go get github.com/lokhman/json-zeroer

## Usage

    package main
    
    import (
        "encoding/json"
        "time"
        
        _ "github.com/lokhman/json-zeroer"
    )
    
    type Object struct {
        X string    `json:"x,omitempty"`
        Y time.Time `json:"y,omitempty"`
        Z time.Time `json:"z,omitempty"`
    }
    
    func main() {
        o := Object{Y: time.Now()}
        data, err := json.Marshal(o)
        if err != nil {
            panic(err)
        }
        println(string(data)) // {"y":"2018-03-15T19:44:50Z"}
    }

See [`examples`](examples/examples.go) for more examples.

## Tests

Use `go test` for testing.

## Notes

Library should work in Linux and Windows OS both x32 and x64. As library uses direct memory access, I would recommend
to use it with caution. Other limitations are listed here: https://github.com/bouk/monkey#notes.

## License

Library is available under the MIT license. The included LICENSE file describes this in detail.
