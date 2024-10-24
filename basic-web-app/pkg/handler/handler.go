package handler

import (
	"fmt"
	"net/http"

	"github.com/Dharmadurai95/go-basic-web-application/pkg/config"
	"github.com/Dharmadurai95/go-basic-web-application/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	err := render.ParseTemplate(w, "home.page.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		fmt.Println("Error parsing home page template:", err)
		return
	}
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	err := render.ParseTemplate(w, "about.page.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		fmt.Println("Error parsing home page template:", err)
		return
	}
}

func addNumber(a, b int) int {
	return a + b
}
