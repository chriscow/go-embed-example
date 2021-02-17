package main

import (
	"log"
	"net/http"
	"webe/web"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	handler := web.AssetHandler("/", "build")
	r.Get("/", handler.ServeHTTP)
	r.Get("/*", handler.ServeHTTP)

	log.Println("Listening on :3333")
	http.ListenAndServe(":3333", r)
}
