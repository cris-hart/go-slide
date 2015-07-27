package animal

type Animal struct {
	Name string
}

func (h *Human) MyNameIs() string {
	return h.Name
}
