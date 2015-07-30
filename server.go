package main

import (
	"flag"
	"log"
	"net/http"
)

var addr = flag.String("addr", ":1718", "http service address")

func main() {
	flag.Parse()
	http.Handle("/", http.HandlerFunc(HI))
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
	log.Printf("Go to http://localhost:%d for a greeting \n", *addr)
}

func HI(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("hello"))
}

type dummyResponse struct {
	what string
}

func (dr dummyResponse) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ciao " + dr.what))
}

func check(fn http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		if len(r.URL.Query().Get("key")) == 0 {
			http.Error(w, "missing key", http.StatusUnauthorized)
			return
		}
		fn(w, r)
	}

}
