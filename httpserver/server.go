package speedchecksystem

import (
	// "log"
	"net/http"
)

// func ServerHTTP() {
	// http.Handle("/datastore", http.FileServer(http.Dir("./datastore/filesystem")))
	// http.ListenAndServe(":8080", nil)
// }

type HttpHandler struct {}

func (h HttpHandler) ServeHTTP(res http.ResponseWriter, req *http.Request){
	data := []byte("Hello Balodya!")

	res.Write(data)
}