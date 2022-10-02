package mock_gomock

type Foo interface {
	Bar(x int) int
}

type User interface {
	Update(admin Admin) error
}

type Admin struct {
	ID uint32
}

func Bar(x int) int {
	return x + 2
}
