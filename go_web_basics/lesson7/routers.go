package main

import (
	"encoding/json"
	"net/http"
	"serv/models"

	"github.com/go-chi/chi"

	httpSwagger "github.com/swaggo/http-swagger"
)

func ConfigureRouters(r *chi.Mux) {
	r.Route("/", func(r chi.Router) {
		r.Route("/api/v1", func(r chi.Router) {
			r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8080/api/v1/docs/swagger.json")))
			r.Get("/docs/swagger.json", HandleSwaggerJSON)

			r.Post("/echo", HandleEcho)
			r.Put("/echo", HandleEcho2)
		})
	})
}

func SendErr(w http.ResponseWriter, code int, err error, obj ...interface{}) {
	e := models.ServErr{
		Code:     code,
		Err:      err.Error(),
		Desc:     "server error",
		Internal: obj,
	}
	data, _ := json.Marshal(e)
	w.Write(data)
}
