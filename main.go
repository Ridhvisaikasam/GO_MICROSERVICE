package main

import (
	"context"
	"go_microservice/data"
	"go_microservice/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
)

// main entry point
func main() {

	l := log.New(os.Stdout, "product-api(logger)", log.LstdFlags)
	v := data.NewValidation()
	//all the content in handle func into an independent object
	//hh := handlers.NewHello(l)
	//gh := handlers.NewGoodbye(l)
	ph := handlers.NewProducts(l, v)

	//new servemux
	//sm := http.NewServeMux()
	//refactoring by using gorilla mux router
	sm := mux.NewRouter()
	//using subrouters for more flexibility and easy usage of middleware
	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", ph.GetProducts)

	putRouter := sm.Methods(http.MethodPut).Subrouter()
	//gorilla automatically pulls out id
	putRouter.HandleFunc("/{id:[0-9]+}", ph.UpdateProducts)
	//goes into handle func only after middleware
	putRouter.Use(ph.MiddlewareProductValidation)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", ph.AddProducts)
	//goes into handle func only after middleware
	postRouter.Use(ph.MiddlewareProductValidation)

	deleteRouter := sm.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/{id:[0-9]+}", ph.DeleteProduct)

	//swagger.yaml must be served as used by Redoc
	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(opts, nil)
	// serve docs
	getRouter.Handle("/docs", sh)
	//go has an inbuilt file server
	getRouter.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	//register handler to servemux ,,, callls the serve http associated with it
	//sm.Handle("/", hh)
	//sm.Handle("/goodbye", gh)
	//sm.Handle("/products/", ph)

	/*timeouts are imp -- resources are finite , if client pauses (blocked conn) ,,,
	multiple blocked connections -- server fails
	service denial attack*/

	//tuning elements by manually creating a server(http)
	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

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
	//http.ListenAndServe(":9090", sm)

	//starting a server
	//refactor a little because of shutdown
	//handling in go func so its not gonna block
	go func() {
		l.Println("Starting server on port 9090")

		err := s.ListenAndServe()
		if err != nil {
			//print followed by os.exit()
			l.Fatal(err)
		}
	}()

	//broadcast a message on a channel whenever os is interrupted or killed
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	//block here until a message available to conceive ,,, once conceived shutdown
	sig := <-sigChan
	l.Println("Received terminate ,graceful shutdown :: ", sig)

	/*graceful shutdown -imp- during large file upload/db transaction of client
	now we need to upgrade the server ,,, if we shutdown we are obstructing clients task
	but with GO server after shutdown() called
	doesnot allow new req ,, wait till all the requests completed ,, program will exit
	*/

	//this implies that allow 30 sec to do all of work and then forcefully close them
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)

}

//good structured packages required for microservices
//doing this deliberately for merge conflicts
//start to refactor , need to bring in better practices and patterns to structure the code
//lot of code in func main , need to think about testing
