package main

import(
	"fmt"
	"encoding/json"
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
	hub.Requests = MakeRequestDatabase(h.maxRequests)
	
	h.hubs[id] = hub
	return hub, nil
}

func (h *HubDatabase) Get(id string) *Hub {
	hub, _ := h.hubs[id]
	return hub
}

func (h *HubDatabase) Delete(id string) {
	delete(h.hubs, id)
}

func (h *HubDatabase) ToJson() ([]byte, error) {
	keys := make([]string, len(h.hubs))

	for _, v := range h.hubs {
		if v.Id != "" {
			keys = append(keys, v.Id)
		}
	}

	return json.Marshal(keys)
}


