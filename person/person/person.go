package person

type Person struct {
	name string
	age  int
}

func (p *Person) Set(name string, age int) {
	p.name = name
	p.age = age
}
