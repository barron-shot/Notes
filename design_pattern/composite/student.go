package composite

import "fmt"

type Student struct {
	name  string
	level string
}

func (s Student) Display() {
	fmt.Printf("name: %v, level: %v \n", s.name, s.level)
}

func NewStudent(name, level string) Student {
	return Student{
		name:  name,
		level: level,
	}
}
