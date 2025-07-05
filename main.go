package main

import (
	"embed"
	"net/http"
	"strings"
)

//go:embed README.md play.sh
var page embed.FS

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// read the markdown file
		data, err := page.ReadFile("README.md")
		if err != nil {
			http.Error(w, "Failed to read README.md", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text/plain")
		w.Write(data)
	})

	mux.HandleFunc("/play", func(w http.ResponseWriter, r *http.Request) {
		data, err := page.ReadFile("play.sh")
		if err != nil {
			http.Error(w, "Failed to read play.sh", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/plain")
		w.Write(data)
	})

	if err := http.ListenAndServe(":8081", mux); err != nil {
		panic(err)
	}
}

func isTerminal(r *http.Request) bool {
	terminals := []string{"curl", "wget", "lynx", "links", "elinks", "httpie"}
	userAgent := r.Header.Get("User-Agent")
	for _, terminal := range terminals {
		if strings.Contains(userAgent, terminal) {
			return true
		}
	}
	return false
}
