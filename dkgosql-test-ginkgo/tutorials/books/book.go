package books

type Book struct {
	Name      string
	Religious bool
}

func (b *Book) IsReligious() bool {
	if b.Religious {
		return true
	}

	return false
}
