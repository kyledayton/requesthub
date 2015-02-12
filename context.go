package main

import "fmt"
import "net/http"
import "log"
import "flag"

const(
	DEFAULT_PORT = 54321
)

func Start() {
	var parseReq = flag.Int("r", DEFAULT_MAX_REQUESTS, "max requests to store")
	var parsePort = flag.Int("p", DEFAULT_PORT, "which port to bind to")
	flag.Parse();

	maxRequests := *parseReq
	port := *parsePort

	log.Printf("Creating in-memory database (max %d)\n", maxRequests)
	db := newHubDatabase(maxRequests)
	db.Create("default")
	
	http.HandleFunc("/requests", func(w http.ResponseWriter, r *http.Request) {
		json, err := db.Get("default").Requests.ToJson()
	
		if err != nil {
			log.Panic(err)
		}

		w.Header().Add("Content-Type", "application/json")
		w.Write(json)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "GET" {
			w.Header().Add("Content-Type", "text/html")
			w.Write([]byte(PAGE_CONTENT))
		} else {
			db.Get("default").Requests.Insert(r)
		}

	})

	http.HandleFunc("/clear", func(w http.ResponseWriter, r *http.Request) {
		db.Get("default").Requests.Clear()
	});

	log.Printf("Listening on port %d", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
