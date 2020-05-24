package demos

type Student struct {
	name   string
	age    uint8
	gender string
}

func NewStudent(name string, age uint8, gender string) *Student {
	return &Student{name: name, age: age, gender: gender}
}

