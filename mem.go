package mem

type Foo interface {
	Bar() string
}

type foo struct {
	a string
}

func (f foo) Bar() string {
	return "Bar" + f.a
}

func printer(f Foo) string {
	return f.Bar()
}
