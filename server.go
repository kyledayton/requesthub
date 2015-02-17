package main

import(
	"fmt"
	"net/http"
	"log"
	"flag"
	"strings"
	"html/template"
)


func Start() {
	var parseReq = flag.Int("r", DEFAULT_MAX_REQUESTS, "max requests to store")
	var parsePort = flag.Int("p", DEFAULT_PORT, "which port to bind to")
	flag.Parse();

	maxRequests := *parseReq
	port := *parsePort

	viewPage := template.Must(template.New("show").Parse(HUB_PAGE_CONTENT))
	indexPage := template.Must(template.New("index").Parse(INDEX_PAGE_CONTENT))

	log.Printf("Initializing hub database (maxReq=%d)\n", maxRequests)
	db := newHubDatabase(maxRequests)

	router := MakeRouter()

	router.HandleFunc(`/([\w\d\-_]+)/forward`, func(w http.ResponseWriter, r *http.Request) {
		parts := strings.Split(r.URL.Path, "/")
		hubName := parts[1]

		if hubName != "" {
			hub := db.Get(hubName)
			dest := strings.TrimSpace( r.FormValue("url") )
			
			if hub != nil {
				hub.ForwardURL = dest
			}
		}
	})

	router.HandleFunc(`/([\w\d\-_]+)/requests`, func(w http.ResponseWriter, r *http.Request) {
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

	router.HandleFunc(`/([\d\w\-_]+)/clear`, func(w http.ResponseWriter, r *http.Request) {
		parts := strings.Split(r.URL.Path, "/")
		hubName := parts[1]
		hub := db.Get(hubName)

		if hub != nil {
			db.Get(hubName).Requests.Clear()
		}
	})

	router.HandleFunc(`/([\d\w\-_]+)`, func(w http.ResponseWriter, r *http.Request) {
		parts := strings.Split(r.URL.Path, "/")
		hubName := parts[1]

		hub := db.Get(hubName)

		if hub != nil {
			if r.Method == "GET" {
				w.Header().Add("Content-Type", "text/html")
				viewPage.Execute(w, hub)
			} else {
				hub.Requests.Insert(r)
			}
		}
	})

	router.HandleFunc(`/`, func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "GET" {
			w.Header().Add("Content-Type", "text/html")
			indexPage.Execute(w, db.hubs)
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
