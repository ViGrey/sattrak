package main

import (
	"io/ioutil"
	"mime"

	//"net"
	"net/http"
	//"os"
	"strings"
	"time"
)

var (
	httpPath = basePath + "/http/"
)

func getMIMEType(path string) (mimeType string) {
	pathSplit := strings.Split(path, ".")
	extension := pathSplit[len(pathSplit)-1]
	if len(path) > 0 {
		mimeType = mime.TypeByExtension("." + extension)
	}
	return
}

func catchAll(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	url := r.URL.Path
	url = url[1:]
	urlModded := strings.ToLower(url)
	if len(urlModded) > 0 {
		if urlModded[len(urlModded)-1] == '/' {
			urlModded = urlModded[:len(urlModded)-1]
		}
	}
	if urlModded == "" {
		urlModded = "index"
	}
	if urlModded == "index" {
		handlePath(w, r, "index.html", httpPath)
	}
	if urlModded == "display" {
		handlePath(w, r, "display.html", httpPath)
	}
	if urlModded == "control" {
		handlePath(w, r, "control.html", httpPath)
	}
	if strings.Index(urlModded, "fonts/") == 0 {
		handlePath(w, r, url, httpPath)
	}
	if strings.Index(urlModded, "scripts/") == 0 {
		handlePath(w, r, url, httpPath)
	}
	if strings.Index(urlModded, "styles/") == 0 {
		handlePath(w, r, url, httpPath)
	}
	if strings.Index(urlModded, "images/") == 0 {
		handlePath(w, r, url, httpPath)
	}
}

func handlePath(w http.ResponseWriter, r *http.Request, path, pathBase string) {
	if content, err := ioutil.ReadFile(pathBase + path); err == nil {
		w.Header().Set("content-type", getMIMEType(path))
		w.Write(content)
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 page not found"))
	}
}

func httpListen() {
	mux := http.NewServeMux()
	mux.HandleFunc("/ws", wsHandler)
	mux.HandleFunc("/", catchAll)
	srv := &http.Server{
		Addr:        "0.0.0.0:1234",
		ReadTimeout: 250 * time.Millisecond,
		Handler:     mux,
	}
	srv.ListenAndServe()
}
