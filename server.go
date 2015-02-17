package main

import(
	"fmt"
	"net/http"
	"log"
	"flag"
	"regexp"
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

	router := MakeRouter()

	router.HandleFunc(regexp.MustCompile(`/requests`), func(w http.ResponseWriter, r *http.Request) {
		json, err := db.Get("default").Requests.ToJson()
	
		if err != nil {
			log.Panic(err)
		}

		w.Header().Add("Content-Type", "application/json")
		w.Write(json)
	})

	router.HandleFunc(regexp.MustCompile(`/hub`), func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/html")
		w.Write([]byte(INDEX_PAGE_CONTENT))
	})

	router.HandleFunc(regexp.MustCompile(`/clear`), func(w http.ResponseWriter, r *http.Request) {
		db.Get("default").Requests.Clear()
	});

	router.HandleFunc(regexp.MustCompile(`/`), func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "GET" {
			w.Header().Add("Content-Type", "text/html")
			w.Write([]byte(HUB_PAGE_CONTENT))
		} else {
			db.Get("default").Requests.Insert(r)
		}

	})

	log.Printf("Listening on port %d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}
