package main

import "fmt"
import "sync"
import "net/http"
import "io/ioutil"

const(
	DEFAULT_MAX_REQUESTS = 256
)

type Hub struct {
	Id string
	Requests *RequestDatabase
	ForwardURL string
}

type HubDatabase struct {
	hubs map[string]*Hub
	maxRequests int
}

func newHubDatabase(maxRequests int) *HubDatabase {
	db := &HubDatabase{make(map[string]*Hub), maxRequests}
	return db
}

func (h *HubDatabase) Create(id string) (*Hub, error) {
	_, exists := h.hubs[id]

	if exists {
		return nil, fmt.Errorf("Hub %s is already in use", id)
	}

	hub := new(Hub)
	hub.Id = id
	hub.Requests = newRequestDatabase(h.maxRequests)
	
	h.hubs[id] = hub
	return hub, nil
}

func (h *HubDatabase) Get(id string) *Hub {
	hub, _ := h.hubs[id]
	return hub
}

type RequestDatabase struct {
	*sync.RWMutex
	requests map[string][]*Request
	maxRequests int
}

type Request struct {
	Header http.Header `json:"headers"`
	ContentLength int64 `json:"content_length"`
	Body string	`json:"body"`
	Method string `json:"method"`
}

func newRequestDatabase(maxRequests int) *RequestDatabase {
	db := &RequestDatabase{new(sync.RWMutex), make(map[string][]*Request), maxRequests}
	db.requests["requests"] = make([]*Request, 0, maxRequests)
	return db
}

func (d *RequestDatabase) Insert(req *http.Request) {
	d.Lock()
		r := cloneRequest(req)
		requests := d.requests["requests"]

		d.requests["requests"] = append([]*Request{r}, requests...)

		if len(d.requests["requests"]) >= d.maxRequests {
			d.requests["requests"] = d.requests["requests"][0:d.maxRequests]
		}

	d.Unlock()
}

func cloneRequest(req *http.Request) *Request {
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

func (d *RequestDatabase) Clear() {
	d.Lock()
		d.requests = make(map[string][]*Request)
		d.requests["requests"] = make([]*Request, 0, d.maxRequests)
	d.Unlock();
}
