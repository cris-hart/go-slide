package animal

type Animal struct {
	Name string
}

func (h *Animal) MyNameIs() string {
	return h.Name
}
