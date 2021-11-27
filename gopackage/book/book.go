package book

type book struct {
	Name string
}

func NewBook() book {
	return book{
		Name: "Harry",
	}
}
