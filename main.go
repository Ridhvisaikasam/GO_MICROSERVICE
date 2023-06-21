package main

import (
	"go_microservice/handlers"
	"log"
	"net/http"
	"os"
)

// main entry point
func main() {

	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	//all the content in handle func into an independent object
	hh := handlers.NewHello(l)
	gh := handlers.NewGoodbye(l)

	//new servemux
	sm := http.NewServeMux()
	//register handler to servemux
	sm.Handle("/", hh)
	sm.Handle("/goodbye", gh)

	//register handler with the server
	//converting function to handle and registering to defaultservemux(server multiplexer for multiple paths)
	//http.HandleFunc()

	//path(when matched or not matches any other),function to be executed which handles http req and res
	/*http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {......
	})
	http.HandleFunc("/bye", func(http.ResponseWriter, *http.Request) {
		log.Println("bye")
	})*/

	//listen at port,serve handler(interface implementing servehttp)(nil->default servemux(redirecting paths))
	http.ListenAndServe(":9090", sm)

}

//good structured packages required for microservices
//doing this deliberately for merge conflicts
//start to refactor , need to bring in better practices and patterns to structure the code
//lot of code in func main , need to think about testing
