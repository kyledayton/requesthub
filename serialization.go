package main

import "encoding/json"

func (d *RequestDatabase) ToJson() ([]byte, error) {
	return json.Marshal(d.requests)
}
