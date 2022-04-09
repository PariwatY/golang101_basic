package customer

type Person struct {
	name string
	age  int
}

func (p Person) GetName() string {
	return p.name
}

func (p *Person) SetName(name string) {
	p.name = name
}
