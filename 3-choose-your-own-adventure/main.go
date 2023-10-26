package main

import (
	"encoding/json"
	"html/template"
	"io"
	"net/http"
	"os"
)

func main() {
	r := http.NewServeMux()

	j, _ := os.Open("gopher.json")
	b, _ := io.ReadAll(j)

	m := make(map[string]struct {
		Title   string   `json:"title"`
		Story   []string `json:"story"`
		Options []struct {
			Text string `json:"text"`
			Arc  string `json:"arc"`
		}
	})
	json.Unmarshal(b, &m)

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		a := r.URL.Query().Get("arc")

		d, e := m[a]
		if !e {
			d = m["intro"]
		}
		if d.Title == "" {
			d = m["intro"]
		}

		t := template.Must(template.ParseFiles("./arc.html"))
		t.Execute(w, d)
	})

	http.ListenAndServe(":4321", r)
}
