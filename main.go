package main

import (
	"os"
	"time"
	"net/http"
)

type handler struct {
}

func (h *handler) ServeHTTP(rw http.ResponseWriter, r *http.Request){
	rw.Write([]byte("Hi man"))
}

func main(){
	port := os.Getenv("WSPORT")

	if port == ""{
		port = ":8080"
	}

	s := &http.Server{
		Addr:           ":8080",
		Handler:        &handler{},
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()

}