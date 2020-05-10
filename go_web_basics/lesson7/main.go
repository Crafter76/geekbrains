package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewMux()
	ConfigureRouters(r)
	log.Printf("running")
	http.ListenAndServe(":8080", r)
}

func StringConcat(a, b string) string {
	return fmt.Sprintf("%s %s", a, b)
}
