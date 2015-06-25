package main

import(
	"fmt"
	"net/http"
	"log"
	"strings"
	"strconv"
	"html/template"

	"github.com/kyledayton/requesthub/templates"
)

var config *Config

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

func modernizrJShandler(w http.ResponseWriter, r *http.Request) {
	f, _ := assets_modernizr_js_bytes()
	w.Header().Set("Content-Type", "application/javascript")
	w.Write(f)
}

func authFailed(r *http.Request) bool {
	reqUser, reqPass, ok := r.BasicAuth()

	return config.AuthEnabled() && (!ok || reqUser != config.Username || reqPass != config.Password)
}

func requireAuth(w http.ResponseWriter) {
	w.Header().Add("WWW-Authenticate", "Basic realm=\"requesthub\"")
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte("Authorization Failed."))
}

func Start() {
	config = NewConfig()

	log.Printf("Initializing hub database (maxReq=%d)\n", config.MaxRequests)
	db := newHubDatabase(config.MaxRequests)

	if config.HasYAMLConfig() {
		err := config.ApplyYAMLConfig(db)
		if err != nil {
			log.Printf("Error loading config file (%s): %s", config.YamlConfigFile, err)
		}
	}

	if config.AuthEnabled() {
		log.Printf("Using HTTP Basic Auth\n")
	}

	viewPage := template.Must(template.New("show").Parse(templates.SHOW_HUB))
	indexPage := template.Must(template.New("index").Parse(templates.INDEX))


	forwardClient := new(http.Client)

	router := MakeRouter()

	router.HandleFunc(`/assets/foundation.css`, foundationCSShandler)
	router.HandleFunc(`/assets/foundation.js`, foundationJShandler)
	router.HandleFunc(`/assets/jquery.js`, jqueryJShandler)
	router.HandleFunc(`/assets/modernizr.js`, modernizrJShandler)

	router.HandleFunc(`/show/([\d\w\-_]+)`, func(w http.ResponseWriter, r *http.Request) {
		if authFailed(r) {
			requireAuth(w)
			return
		}

		parts := strings.Split(r.URL.Path, "/")
		hubName := parts[2]
		hub := db.Get(hubName)

		if hub != nil {
			w.Header().Add("Content-Type", "text/html")

			if authFailed(r) {
				requireAuth(w)
				return
			}

			viewPage.Execute(w, hub)
		}
	})

	router.HandleFunc(`/([\w\d\-_]+)/forward`, func(w http.ResponseWriter, r *http.Request) {
		if authFailed(r) {
			requireAuth(w)
			return
		}

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
		if authFailed(r) {
			requireAuth(w)
			return
		}

		parts := strings.Split(r.URL.Path, "/")
		hubName := parts[1]

		if hubName != "" {
			hub := db.Get(hubName)

			if hub != nil {
				w.Write([]byte(strconv.Itoa(int(hub.Requests.Count))))
			}
		}
	})

	router.HandleFunc(`/([\w\d\-_]+)/delete`, func(w http.ResponseWriter, r *http.Request) {
		if authFailed(r) {
			requireAuth(w)
			return
		}

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
		if authFailed(r) {
			requireAuth(w)
			return
		}

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
		if authFailed(r) {
			requireAuth(w)
			return
		}

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
			req := hub.Requests.Insert(r)

			if hub.ForwardURL != "" {
				go req.Forward(forwardClient, hub.ForwardURL)
			}
		}
	})


	router.HandleFunc(`/`, func(w http.ResponseWriter, r *http.Request) {

		if authFailed(r) {
			requireAuth(w)
			return
		}

		if r.Method == "GET" {
			w.Header().Add("Content-Type", "text/html")
			indexPage.Execute(w, db.hubs)
		} else if r.Method == "POST" {
			hubName := strings.TrimSpace( r.FormValue("hub_name") )

			if hubName == "" {
				http.Redirect(w, r, "/", 302)
				return
			}

			hub, err := db.Create(hubName)

			if err != nil {
				http.Redirect(w, r, `/`, 302)
			} else {
				http.Redirect(w, r, fmt.Sprintf("/show/%s", hub.Id), 302)
			}
		}
	})

	log.Printf("Listening on port %d\n", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), router))
}
