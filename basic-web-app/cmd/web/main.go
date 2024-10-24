package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Dharmadurai95/go-basic-web-application/pkg/config"
	"github.com/Dharmadurai95/go-basic-web-application/pkg/handler"
	"github.com/Dharmadurai95/go-basic-web-application/pkg/render"
)

const PORT = ":7000"

func main() {
	var a config.AppConfig
	a.UseCache = true
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Could not able to create template cache")
	}
	a.TemplateCache = tc
	newRepo := handler.NewRepo(&a)
	handler.NewHandlers(newRepo)

	render.NewTemplage(&a)

	server := &http.Server{
		Addr:    PORT,
		Handler: Routes(&a),
	}
	servererr := server.ListenAndServe()
	if servererr != nil {
		fmt.Println("Error starting server:", err)
	}
}
