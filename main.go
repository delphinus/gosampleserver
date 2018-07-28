package main

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {
	if err := process(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func process() error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var in io.ReadCloser
		var contentType string
		var status int
		var err error
		switch r.Method {
		case "GET":
			in, contentType, status, err = get(r)
		case "POST":
			status, err = post(r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		if err != nil {
			w.WriteHeader(status)
			_, _ = w.Write([]byte(fmt.Sprintf("%v", err)))
			return
		}
		var body []byte
		if in != nil {
			defer in.Close()
			body, err = ioutil.ReadAll(in)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				_, _ = w.Write([]byte(fmt.Sprintf("%v", err)))
				return
			}
			w.Header().Add("Content-Length", strconv.Itoa(len(body)))
		}
		w.Header().Add("Content-Type", contentType)
		w.WriteHeader(status)
		_, _ = w.Write(body)
	})
	return http.ListenAndServe(":8080", nil)
}

func get(r *http.Request) (
	in io.ReadCloser, contentType string, status int, err error) {
	path := r.URL.Path
	if path == "/" {
		path = "/index.html"
	}
	isFile := true
	switch {
	case strings.HasSuffix(path, ".html"):
		contentType = "text/html"
	case strings.HasSuffix(path, ".js"):
		contentType = "application/javascript"
	case path == "/counter.json":
		isFile = false
		contentType = "applicatjon/json"
	default:
		contentType = "application/octet-stream"
	}
	if isFile {
		in, status, err = readFile(path)
	} else {
		in, status, err = getCounterJSON()
	}
	return
}

func post(r *http.Request) (int, error) {
	if r.URL.Path != "/counter" {
		return http.StatusNotFound, errors.New("path not found")
	}
	if r.Body == nil {
		return http.StatusBadRequest, errors.New("input not found")
	}
	defer r.Body.Close()
	return putCounterJSON(r.Body)
}
