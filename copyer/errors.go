package copyer

type Error struct {
	err string
}

func (e *Error) Error() string {
	return e.err
}
