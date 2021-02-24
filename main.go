package main

import (
	"flag"
	"fmt"
	// "log"
	"net/http"
	"time"
	"github/Stanlyzoolo/speedchecksystem/httpserver"
	
)

func main() {

	t := time.Now()
	fmt.Println(t.Format("02.01.2006 15:04:05"))

	// log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("/datastore"))))

	// port := flag.String("p", "8100", "port to serve on")
	directory := flag.String("d","./datastore/", "the directory of static file to host")
	flag.Parse()

	http.Handle("/", http.FileServer(http.Dir(*directory)))

	// log.Printf("Serving %s on HTTP port: %s\n", *directory, *port)
	// log.Fatal(http.ListenAndServe(":"+*port, nil))

	handler := httpserver.HttpHanler{}
	http.ListenAndServe(":9000", handler)
}
