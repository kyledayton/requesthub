package main

import(
	"net/http"
	"sync"
	"io/ioutil"
	"encoding/json")

type Request struct {
	Header http.Header `json:"headers"`
	ContentLength int64 `json:"content_length"`
	Body string	`json:"body"`
	Method string `json:"method"`
}


type RequestDatabase struct {
	*sync.RWMutex
	requests []*Request
	maxRequests int
}

func MakeRequest(req *http.Request) *Request {
	r := new(Request)

	r.Header = make(http.Header)

	for k, v := range req.Header {
		r.Header[k] = v
	}
	
	r.ContentLength = req.ContentLength
	r.Method = req.Method

	body, _ := ioutil.ReadAll(req.Body)
	r.Body = string(body)

	return r
}

func (r *Request) ToJson() ([]byte, error) {
	return json.Marshal(r)
}

/////////////////////////////
// Request Database
/////////////////////////////

func (d *RequestDatabase) Insert(req *http.Request) {
	d.Lock()
		r := MakeRequest(req)
		d.requests = append([]*Request{r}, d.requests...)

		if len(d.requests) >= d.maxRequests {
			d.requests = d.requests[0:d.maxRequests]
		}

	d.Unlock()
}

func (d *RequestDatabase) Clear() {
	d.Lock()
		d.requests = make([]*Request, 0, d.maxRequests)
	d.Unlock();
}

func (d *RequestDatabase) ToJson() ([]byte, error) {
	return json.Marshal(d.requests)
}

func MakeRequestDatabase(capacity int) *RequestDatabase {
	db := &RequestDatabase{new(sync.RWMutex), make([]*Request, 0, capacity), capacity}
	return db
}
