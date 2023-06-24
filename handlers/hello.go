package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//handler is an interface with just serveHTTP method with rw and req

// something related to dependancy injection
type Hello struct {
	l *log.Logger
}

// in test we can just replace logger to anything (configuring in broader level)
func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

// add method to satisfy interface
func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	//dont create concrete objects inside handlers like loggers or db
	h.l.Println("hello")
	//reading the body from req
	d, err := ioutil.ReadAll(r.Body)
	//error handling
	if err != nil {
		http.Error(rw, "oops", http.StatusBadRequest)
		//rw.WriteHeader(http.StatusBadRequest)
		//rw.Write([]byte("oops"))
		return
	}

	fmt.Fprintf(rw, "Data %s \n", d)
}
