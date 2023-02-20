package main

import (
	"bytes"
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type Entry struct {
	Reachable bool   `json:"reachable"`
	Status    int    `json:"status"`
	DateTime  string `json:"dateTime"`
}

type TemplateData struct {
	CheckUrl string
	Entries  []Entry
}

func getStatus(checkUrl string) Entry {
	now := time.Now()
	res, err := http.Get(checkUrl)
	if err != nil {
		return Entry{false, -1, now.Format(time.RFC3339)}
	}

	return Entry{true, res.StatusCode, now.Format(time.RFC3339)}
}

func getHistory(isInitialBuild bool, selfUrl string) []Entry {
	if isInitialBuild {
		return []Entry{}
	}

	res, err := http.Get(selfUrl + "/history.json")
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()

	entries := []Entry{}

	err = json.NewDecoder(res.Body).Decode(&entries)
	if err != nil {
		log.Fatalln(err)
	}

	return entries
}

func main() {
	// Get config
	isInitialBuild := os.Getenv("INITIAL_BUILD") == "true"
	checkUrl := os.Getenv("CHECK_URL")
	selfUrl := os.Getenv("SELF_URL")

	// Validate config
	if checkUrl == "" {
		log.Fatalln("No check URL configured")
	}

	if selfUrl == "" && !isInitialBuild {
		log.Fatalln("No self URL configured")
	}

	// Get current status and history
	status := getStatus(checkUrl)
	history := getHistory(isInitialBuild, selfUrl)
	entries := append(history, status)

	// Write history
	raw, err := json.Marshal(entries)
	if err != nil {
		log.Fatalln(err)
	}
	ioutil.WriteFile(filepath.Join("static", "history.json"), []byte(raw), 0644)

	// Load template
	template, err := template.ParseFiles(filepath.Join("src", "template.html"))
	if err != nil {
		log.Fatalln(err)
	}

	// Render template
	var tpl bytes.Buffer
	if err := template.Execute(&tpl, TemplateData{checkUrl, entries}); err != nil {
		log.Fatalln(err)
	}

	// Write index.html
	result := tpl.String()
	err = ioutil.WriteFile(filepath.Join("static", "index.html"), []byte(result), 0644)
	if err := template.Execute(&tpl, TemplateData{checkUrl, entries}); err != nil {
		log.Fatalln(err)
	}
}
