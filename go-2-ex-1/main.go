package main

import "fmt"

type FullName struct {
	// TODO: add fields
	FirstName string
	LastName  string
}
// TODO: declare a structure for birth date
type BirthDate struct {
	Day int
	Month int
	Year int
}


type Profile struct {
	Name             FullName
	BirthDate        BirthDate
	NumberOfSiblings byte
	ZodiacSign       rune
}

func main() {
	
	var me = Profile{
		// TODO: set name and birth date information
		Name: FullName{
			FirstName: "J",
			LastName:  "T",
		},
		BirthDate: BirthDate{
			Day:   1,
			Month: 10,
			Year:  2005,
		},
		NumberOfSiblings: 3,   // TODO: adjust
		ZodiacSign:       'W', // TODO: adjust
	}
	fmt.Println(me)

	fmt.Println("Siblings Before:", me.NumberOfSiblings)
	// TODO: imagine, you get a little brother or sister
	me.NumberOfSiblings++
	fmt.Println("Siblings After:", me.NumberOfSiblings)
}
