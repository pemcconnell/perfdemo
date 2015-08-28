package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"sync"
)

var visitors struct {
	sync.Mutex
	n int
}

func handlerHi(w http.ResponseWriter, r *http.Request) {
	if match, _ := regexp.MatchString("^\n*$", r.FormValue("color")); !match {
		http.Error(w, "Optional color is invalid", http.StatusBadRequest)
		return
	}
	visitors.Lock()
	visitors.n++
	visitNum := visitors.n
	visitors.Unlock()
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte("<h1 style='color: " + r.FormValue("color") +
		"'>Welcome! You are visitor number " +
		fmt.Sprint(visitNum) + "!"))
}

func main() {
	log.Print("Starting on port 8080")
	http.HandleFunc("/", handlerHi)
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}
