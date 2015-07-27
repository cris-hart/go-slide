package "human"

type Human struct {
	Name string
	Phrase string
}

func (h *Human) MyNameIs() string {
	return h.Name
}

func (h *Human) Speak() string {
	return h.Pharese
}