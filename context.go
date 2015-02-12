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
	db := newDatabase(maxRequests)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "GET" {
			json, err := db.ToJson()

			if err != nil {
				log.Panic(err)
			}

			w.Write(json)
		} else {
			db.Insert(r)
		}

	})

	http.HandleFunc("/clear", func(w http.ResponseWriter, r *http.Request) {
		db.Clear()
	});

	log.Printf("Listening on port %d", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
