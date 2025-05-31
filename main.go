package main

import (
	"net/http"
	"strings"
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
		sb.WriteString("\n```\n")

		sb.WriteString("\n\n# Welcome to the ASCII Arcade!\n\n")
		sb.WriteString("## Available Games:\n\n")
		sb.WriteString("- Farkle: `ssh ascii-arcade.games -p 2022`\n")

		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte(sb.String()))
	})

	http.ListenAndServe(":81", mux)
}
