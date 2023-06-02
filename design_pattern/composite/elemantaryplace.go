package composite

func ShowAllStudent() {

	l5student := NewStudent("Paul", "l5")
	l3studentMary := NewStudent("Mary", "l3")
	l3studentLee := NewStudent("Lee", "l3")
	l1student := NewStudent("Barron", "l1")

	l5Level := NewComposite(NewLevel("Level 5"))
	l3Level := NewComposite(NewLevel("Level 3"))
	l1Level := NewComposite(NewLevel("Level 1"))

	levelOrder := NewComposite(l5Level)
	levelOrder.Add(l3Level)
	levelOrder.Add(l1Level)

	l5Level.Add(l5student)
	l3Level.Add(l3studentLee)
	l3Level.Add(l3studentMary)
	l1Level.Add(l1student)

	levelOrder.Display()
	// fmt.Println("------- L5")
	// l5Level.Display()
	// fmt.Println("------- L3")
	// l3Level.Display()
	// fmt.Println("------- L1")
	// l1Level.Display()

}
