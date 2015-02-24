package main

import(
	"fmt"
	"net/http"
	"log"
	"flag"
	"strings"
	"strconv"
	"html/template"

	"github.com/kyledayton/requesthub/templates"
)

func foundationCSShandler(w http.ResponseWriter, r *http.Request) {
	f, _ := assets_foundation_css_bytes()
	w.Header().Set("Content-Type", "text/css")
	w.Write(f)
}

func foundationJShandler(w http.ResponseWriter, r *http.Request) {
	f, _ := assets_foundation_js_bytes()
	w.Header().Set("Content-Type", "application/javascript")
	w.Write(f)
}

func jqueryJShandler(w http.ResponseWriter, r *http.Request) {
	f, _ := assets_jquery_js_bytes()
	w.Header().Set("Content-Type", "application/javascript")
	w.Write(f)
}

func Start() {
	var parseReq = flag.Int("r", DEFAULT_MAX_REQUESTS, "max requests to store")
	var parsePort = flag.Int("p", DEFAULT_PORT, "which port to bind to")
	flag.Parse();

	maxRequests := *parseReq
	port := *parsePort

	viewPage := template.Must(template.New("show").Parse(templates.SHOW_HUB))
	indexPage := template.Must(template.New("index").Parse(templates.INDEX))
	log.Printf("Initializing hub database (maxReq=%d)\n", maxRequests)
	db := newHubDatabase(maxRequests)

	forwardClient := new(http.Client)

	router := MakeRouter()

	router.HandleFunc(`/assets/foundation.css`, foundationCSShandler)
	router.HandleFunc(`/assets/foundation.js`, foundationJShandler)
	router.HandleFunc(`/assets/jquery.js`, jqueryJShandler)

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

	router.HandleFunc(`/([\w\d\-_]+)/latest`, func(w http.ResponseWriter, r *http.Request) {
		parts := strings.Split(r.URL.Path, "/")
		hubName := parts[1]

		if hubName != "" {
			hub := db.Get(hubName)

			if hub != nil {
				w.Write([]byte(strconv.Itoa(int(hub.Requests.lastUpdate.Unix()))))
			}
		}
	})

	router.HandleFunc(`/([\w\d\-_]+)/delete`, func(w http.ResponseWriter, r *http.Request) {
		parts := strings.Split(r.URL.Path, "/")
		hubName := parts[1]

		if hubName != "" {
			hub := db.Get(hubName)
			
			if hub != nil {
				db.Delete(hub.Id)
				http.Redirect(w, r, "/", 302)
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
			hub.Requests.Clear()
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
				req := hub.Requests.Insert(r)
				
				if hub.ForwardURL != "" {
					go req.Forward(forwardClient, hub.ForwardURL) 
				}
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
			}

			http.Redirect(w, r, fmt.Sprintf("/%s", hubName), 302)
		}
	})

	log.Printf("Listening on port %d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}
