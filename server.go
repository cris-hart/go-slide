package main

import (
	"flag"
	"log"
	"net/http"
)

var addr = flag.String("addr", ":1718", "http service address") // Q=17, R=18

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
