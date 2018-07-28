package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/bradfitz/gomemcache/memcache"
)

const (
	memd = "memd:11211"
	key  = "counter"
)

// Counter is a JSON struct
type Counter struct {
	Num int `json:"num"`
}

func getCounterJSON() (io.Reader, int, error) {
	mc := memcache.New(memd)
	item, err := mc.Get(key)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	num, err := strconv.Atoi(string(item.Value))
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	c := Counter{num}
	out := bytes.NewBuffer(nil)
	if err := json.NewEncoder(out).Encode(&c); err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return out, http.StatusOK, nil
}

func putCounterJSON(in io.Reader) (int, error) {
	var c Counter
	if err := json.NewDecoder(in).Decode(&c); err != nil {
		return http.StatusBadRequest, err
	}
	mc := memcache.New(memd)
	item := memcache.Item{
		Key:   key,
		Value: []byte(strconv.Itoa(c.Num)),
	}
	if err := mc.Set(&item); err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}
