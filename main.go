package main

import (
	"github.com/didikprabowo/rajaongkir-go/common"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	r := common.AppRegister()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	srv := &http.Server{
		Handler:      r,
		Addr:         ":" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Println(srv.ListenAndServe())
}
