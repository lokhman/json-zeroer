package json_zeroer

import (
	"encoding/json"
	"testing"
	"time"
)

type Object struct {
	X string    `json:"x,omitempty"`
	Y time.Time `json:"y,omitempty"`
	Z time.Time `json:"z,omitempty"`
}

func Test(t *testing.T) {
	o := Object{Y: time.Unix(1521143090, 0)}
	data, err := json.Marshal(o)
	if err != nil {
		t.Fatal(err)
	}
	if string(data) != `{"y":"2018-03-15T19:44:50Z"}` {
		t.FailNow()
	}
}
