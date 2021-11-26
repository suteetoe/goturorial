package exrepo

import "testing"

type fakeRepository struct{}

func (fakeRepository) QueryLang(id uint) Language {

	lang := Language{
		ID:   1,
		Name: "5555",
	}
	return lang
}

func TestExecute(t *testing.T) {

	h := NewHandler(fakeRepository{})

	s := h.Do(0)
	want := "5555 language"
	if s != want {
		t.Errorf(" want %q bug get %s", want, s)
	}
}
