package main

import "encoding/json"

func (r *Request) ToJson() ([]byte, error) {
	return json.Marshal(r)
}

func (d *Database) ToJson() ([]byte, error) {
	return json.Marshal(d.requests)
}