package main

import "net/http"
import "log"
import "flag"

func Start() {
	var parseReq = flag.Int("r", DEFAULT_MAX_REQUESTS, "max requests to store")
	flag.Parse();

	maxRequests := *parseReq
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

	log.Printf("Listening on port 54321")
	http.ListenAndServe(":54321", nil)
}
