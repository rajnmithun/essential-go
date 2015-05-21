package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var port int

func main() {
	flag.IntVar(&port, "port", 3000, "The port to run the file server on")
	flag.Parse()

	fmt.Printf("Serving the files on localhost:%v\n", port)
	error := ServeStatic(port)

	if error != nil {
		log.Fatalln(error)
	}

}

func ServeStatic(port int) error {
	host := fmt.Sprintf("localhost:%v", port)
	return http.ListenAndServe(host, http.FileServer(http.Dir(".")))
}
