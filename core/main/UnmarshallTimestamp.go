package main

//json unmarshal time that isn't in RFC 3339 format

import (
	"encoding/json"
	"fmt"
	"time"
)

type CustomTime struct {
	time.Time
}

const ctLayout = "2006/01/02|15:04:05"

func (ct *CustomTime) UnmarshalJSON(b []byte) (err error) {
	if b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
	}
	ct.Time, err = time.Parse(ctLayout, string(b))
	return
}

func (ct *CustomTime) MarshalJSON() ([]byte, error) {
	return []byte(ct.Time.Format(ctLayout)), nil
}

type Args struct {
	Time CustomTime
}

var data = `
	{"Time": "2014/08/01|11:27:18"}
`

func main() {
	a := Args{}
	fmt.Println(json.Unmarshal([]byte(data), &a))
	fmt.Println(a.Time.String())
}
