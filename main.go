package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// main entry point
func main() {

	//path(when matched or not matches any other),function to be executed which handles http req and res
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("hello")
		d, err := ioutil.ReadAll(r.Body)
		//error handling
		if err != nil {
			http.Error(rw, "oops", http.StatusBadRequest)
			//rw.WriteHeader(http.StatusBadRequest)
			//rw.Write([]byte("oops"))
			return
		}

		fmt.Fprintf(rw, "Data %s \n", d)
	})

	http.HandleFunc("/bye", func(http.ResponseWriter, *http.Request) {
		log.Println("bye")
	})

	//listen at port,serve handler(interface implementing servehttp)(nil->default servemux(redirecting paths))
	http.ListenAndServe(":9090", nil)

}

//good structured packages required for microservices
//doing this deliberately for merge conflicts
