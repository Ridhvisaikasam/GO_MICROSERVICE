package handlers

import (
	"log"
	"net/http"
)

type Goodbye struct {
	l *log.Logger
}

func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

func (g *Goodbye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	//writes to the response taken by client
	rw.Write([]byte("Byeee"))

	//logger on the server side itself
	g.l.Printf("byeeeeee")
}
