package exrepo

import "fmt"

func Execute() {
	repo := NewStaticRepository()
}

type Language struct {
	ID   uint
	Name string
}

type Handler struct {
	repo Repository
}

type Repository interface {
	QueryLang(uint) Language
}

func NewHandler(repo Repository) *Handler {
	handler := &Handler{
		repo: repo,
	}
	return handler
}

func (h *Handler) Do(id uint) string {
	lang := h.repo.QueryLang(id)
	return fmt.Sprintf("%s language", lang.Name)
}

func NewStaticRepository() *Repository {

}

var data = map[uint]Language{
	1: {ID: 1, Name: "Go"},
	2: {ID: 2, Name: "Java"},
	3: {ID: 3, Name: "Python"},
	4: {ID: 4, Name: "Rust"},
}
