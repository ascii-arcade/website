package main

import (
	"net/http"
	"strings"

	"github.com/russross/blackfriday/v2"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		sb := strings.Builder{}
		sb.WriteString("```\n")
		sb.WriteString(`++------------------------------------------------------------------------------++
++------------------------------------------------------------------------------++
||                                                                              ||
||                                                                              ||
||      _    ____   ____ ___ ___        _    ____   ____    _    ____  _____    ||
||     / \  / ___| / ___|_ _|_ _|      / \  |  _ \ / ___|  / \  |  _ \| ____|   ||
||    / _ \ \___ \| |    | | | |_____ / _ \ | |_) | |     / _ \ | | | |  _|     ||
||   / ___ \ ___) | |___ | | | |_____/ ___ \|  _ <| |___ / ___ \| |_| | |___    ||
||  /_/   \_\____/ \____|___|___|   /_/   \_\_| \_\\____/_/   \_\____/|_____|   ||
||                                                                              ||
||                                                                              ||
||                                                                              ||
++------------------------------------------------------------------------------++
++------------------------------------------------------------------------------++`)
		sb.WriteString("\n```")

		sb.WriteString("\n\n# Welcome to the ASCII Arcade!")
		sb.WriteString("\n\n## Available Games:\n")
		sb.WriteString("- Farkle: `ssh ascii-arcade.games -p 2022`\n")

		if isTerminal(r) {
			w.Header().Set("Content-Type", "text/plain")
			w.Write([]byte(sb.String()))
			return
		}

		// convert markdown to HTML
		htmlFlags := blackfriday.UseXHTML | blackfriday.CompletePage | blackfriday.Smartypants | blackfriday.SmartypantsFractions | blackfriday.SmartypantsDashes | blackfriday.SmartypantsLatexDashes
		renderer := blackfriday.NewHTMLRenderer(blackfriday.HTMLRendererParameters{
			Flags: htmlFlags,
			Title: "ASCII Arcade",
			CSS:   "https://gist.githubusercontent.com/koo04/16ef1dd45e9db08de7e15c94edee2aa3/raw/2e0fa0cfaed99ddef65de1386efe8d1f6b80753e/ascii-games.css",
		})
		output := blackfriday.Run([]byte(sb.String()), blackfriday.WithRenderer(renderer))
		w.Header().Set("Content-Type", "text/html")
		w.Write(output)
	})

	http.ListenAndServe(":81", mux)
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
