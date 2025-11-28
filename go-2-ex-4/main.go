package main

import (
	"fmt"
)

type Student struct {
	FirstName string
	LastName  string
}

type Class struct {
	Name     string
	Students []Student
}

func main() {
	classA := Class{
		Name: "Class A",
		Students: []Student{
			{"Anna", "Müller"},
			{"Ben", "Schneider"},
			{"Clara", "Huber"},
		},
	}

	classB := Class{
		Name: "Class B",
		Students: []Student{
			{"David", "Keller"},
			{"Eva", "Maier"},
			{"Felix", "Hartmann"},
		},
	}

	modules := map[int][]Class{
		346: {classA, classB},
		101: {classA},
		202: {classB},
	}

	fmt.Println("📘 Klassen:")
	fmt.Println(classA)
	fmt.Println(classB)

	fmt.Println("\n Module:")
	for moduleNumber, classes := range modules {
		fmt.Printf("Modul %d wird besucht von:\n", moduleNumber)
		for _, class := range classes {
			fmt.Printf("  %s mit Schülern:\n", class.Name)
			for _, student := range class.Students {
				fmt.Printf("    - %s %s\n", student.FirstName, student.LastName)
			}
		}
	}
}
