package exrepo

import "fmt"

func Execute() {
	repo := NewStaticRepository()
	h := NewHandler(repo)
	s := h.Do(2)
	fmt.Println(s)
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

func NewStaticRepository() Repository {
	repo := &StaticRepo{
		data: data,
	}
	return repo
}

var data = map[uint]Language{
	1: {ID: 1, Name: "Go"},
	2: {ID: 2, Name: "Java"},
	3: {ID: 3, Name: "Python"},
	4: {ID: 4, Name: "Rust"},
}

type StaticRepo struct {
	data map[uint]Language
}

func (s StaticRepo) QueryLang(id uint) Language {
	return s.data[id]
}
