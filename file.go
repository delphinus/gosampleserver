package main

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func readFile(path string) (io.Reader, int, error) {
	f, err := os.Open(filepath.Join("assets", path))
	if err != nil {
		if os.IsNotExist(err) {
			return nil, http.StatusNotFound, err
		}
		return nil, http.StatusInternalServerError, err
	}
	return f, http.StatusOK, nil
}
