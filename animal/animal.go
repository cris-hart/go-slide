package animal

type Animal struct {
	Name string
}

func (h *Animal) MyNameIs() string {
	return h.Name
}

type Parrot struct {
	Animal
	Phrase string
}

func (p *Parrot) Speak() string {
	return p.Phrase
}
