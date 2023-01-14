package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"os"
)

// var tmplt *template.Template
var storyContent map[string]interface{}

/* type Story struct {
	Intro struct {
		Title   string   `json:"title"`
		Story   []string `json:"story"`
		Options []struct {
			Text string `json:"text"`
			Arc  string `json:"arc"`
		} `json:"options"`
	} `json:"intro"`
	NewYork struct {
		Title   string   `json:"title"`
		Story   []string `json:"story"`
		Options []struct {
			Text string `json:"text"`
			Arc  string `json:"arc"`
		} `json:"options"`
	} `json:"new-york"`
	Debate struct {
		Title   string   `json:"title"`
		Story   []string `json:"story"`
		Options []struct {
			Text string `json:"text"`
			Arc  string `json:"arc"`
		} `json:"options"`
	} `json:"debate"`
	SeanKelly struct {
		Title   string   `json:"title"`
		Story   []string `json:"story"`
		Options []struct {
			Text string `json:"text"`
			Arc  string `json:"arc"`
		} `json:"options"`
	} `json:"sean-kelly"`
	MarkBates struct {
		Title   string   `json:"title"`
		Story   []string `json:"story"`
		Options []struct {
			Text string `json:"text"`
			Arc  string `json:"arc"`
		} `json:"options"`
	} `json:"mark-bates"`
	Denver struct {
		Title   string   `json:"title"`
		Story   []string `json:"story"`
		Options []struct {
			Text string `json:"text"`
			Arc  string `json:"arc"`
		} `json:"options"`
	} `json:"denver"`
	Home struct {
		Title   string        `json:"title"`
		Story   []string      `json:"story"`
		Options []interface{} `json:"options"`
	} `json:"home"`
} */

func pathValidation(path string) bool {
	// Check if story contains such key as path provided
	if _, isMapContainsKey := storyContent[path[1:]]; isMapContainsKey {
		return true
	}
	return false
}

func serveAdventure(w http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {

		tmplt, err := template.ParseFiles("index.html")
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte("Something went wrong"))
		}

		if request.URL.Path == "/" {
			story := storyContent["intro"]

			err = tmplt.Execute(w, story)
			if err != nil {
				w.WriteHeader(500)
				w.Write([]byte("Something went wrong"))
			}
		} else if pathValidation(request.URL.Path) {
			story := storyContent[request.URL.Path[1:]]

			err = tmplt.Execute(w, story)
			if err != nil {
				w.WriteHeader(500)
				w.Write([]byte("Something went wrong"))
			}
		} else {
			w.WriteHeader(400)
			w.Write([]byte("Bad Request"))
		}
	}
}


func runServer() {
	http.HandleFunc("/", serveAdventure)

	err := http.ListenAndServe(":8080", nil)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("Server closed\n")
	} else if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
		os.Exit(1)
	}
}

func init() {
	// Just for this specific exercise this implementation is feasible
	// Save story fully to memory as a Map
	file, err := os.ReadFile("gopher.json")
	if err != nil {
		fmt.Println("Error with reading the file")
	}
	json.Unmarshal(file, &storyContent)
}

func main() {
	runServer()
}
