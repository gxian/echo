package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	addr string
	path string
)

func init() {
	flag.StringVar(&addr, "addr", ":8000", "addr=:8000")
	flag.StringVar(&path, "path", "/", "path=/xxx")
}

func main() {
	flag.Parse()
	m := http.NewServeMux()
	m.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		log.Print(r.Header)
		w.WriteHeader(http.StatusOK)
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return
		}
		w.Write(b)
	})
	http.ListenAndServe(addr, m)
}
