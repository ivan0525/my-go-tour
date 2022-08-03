package main

import "fmt"

type Person struct {
	id   int
	name string
}

func showPerson1(person *Person) {
	person.id = 101
	person.name = "Jack"

}

func showPerson2(person Person) {
	person.id = 101
	person.name = "Jack"
	fmt.Printf("person: %v\n", person)
}

func main() {
	person := Person{
		id:   100,
		name: "Tom",
	}

	fmt.Printf("person: %v\n", person)
	showPerson2(person)
	fmt.Printf("person: %v\n", person)
}
