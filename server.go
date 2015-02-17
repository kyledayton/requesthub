package main

import(
	"fmt"
	"net/http"
	"log"
	"flag"
	"regexp"
	"strings"
)


func Start() {
	var parseReq = flag.Int("r", DEFAULT_MAX_REQUESTS, "max requests to store")
	var parsePort = flag.Int("p", DEFAULT_PORT, "which port to bind to")
	flag.Parse();

	maxRequests := *parseReq
	port := *parsePort

	log.Printf("Initializing hub database (maxReq=%d)\n", maxRequests)
	db := newHubDatabase(maxRequests)

	router := MakeRouter()

	router.HandleFunc(regexp.MustCompile(`/([\w\d\-_]+)/requests`), func(w http.ResponseWriter, r *http.Request) {
		parts := strings.Split(r.URL.Path, "/")
		hubName := parts[1]

		if hubName != "" {
			hub := db.Get(hubName)

			if hub != nil {
				json, err := db.Get(hubName).Requests.ToJson()
	
				if err != nil {
					log.Panic(err)
				}

				w.Header().Add("Content-Type", "application/json")
				w.Write(json)
			}
		}
	})

	router.HandleFunc(regexp.MustCompile(`/hubs`), func(w http.ResponseWriter, r *http.Request) {
		json, err := db.ToJson()
		
		if err != nil {
			log.Panic(err)
		}

		w.Header().Add("Content-Type", "application/json")
		w.Write(json)
	})

	router.HandleFunc(regexp.MustCompile(`/([\d\w\-_]+)/clear`), func(w http.ResponseWriter, r *http.Request) {
		parts := strings.Split(r.URL.Path, "/")
		hubName := parts[1]
		hub := db.Get(hubName)

		if hub != nil {
			db.Get(hubName).Requests.Clear()
		}
	})

	router.HandleFunc(regexp.MustCompile(`/([\d\w\-_]+)`), func(w http.ResponseWriter, r *http.Request) {
		parts := strings.Split(r.URL.Path, "/")
		hubName := parts[1]

		if r.Method == "GET" {
			w.Header().Add("Content-Type", "text/html")
			w.Write([]byte(HUB_PAGE_CONTENT))
		} else {
			hub := db.Get(hubName)
			
			if hub != nil {
				hub.Requests.Insert(r)
			}
		}
	})

	router.HandleFunc(regexp.MustCompile(`/`), func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "GET" {
			w.Header().Add("Content-Type", "text/html")
			w.Write([]byte(INDEX_PAGE_CONTENT))
		} else if r.Method == "POST" {
			hubName := strings.TrimSpace( r.FormValue("hub_name") )

			if hubName == "" {
				http.Redirect(w, r, "/", 302)
				return
			}
			
			if db.Get(hubName) == nil {
				db.Create(hubName)
				log.Printf("Created hub %s\n", hubName)
			}

			http.Redirect(w, r, fmt.Sprintf("/%s", hubName), 302)
		}
	})

	log.Printf("Listening on port %d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}
