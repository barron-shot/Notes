package composite

import "fmt"

type Level struct {
	title string
}

func (l Level) Display() {
	fmt.Printf("level: %v\n", l.title)
}

func NewLevel(title string) Level {
	return Level{
		title: title,
	}
}
