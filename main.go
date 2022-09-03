package main

import (
	"fmt"

	"github.com/Asideron/go-two-pointers/set"
)

type Person struct {
	name  string
	knows *set.Set[string]
}

func NewPerson(name string, knows []string) *Person {
	return &Person{
		name:  name,
		knows: set.NewSetFromSlice(knows),
	}
}

func (p *Person) Knows(name string) bool {
	return p.knows.Contains(name)
}

func main() {
	persons := []*Person{
		NewPerson("Alex", []string{"Donald", "Anna", "Bob"}),
		NewPerson("Anna", []string{"Donald", "Bob"}),
		NewPerson("Bob", []string{"Donald", "Alex"}),
		NewPerson("Donald", []string{}),
		NewPerson("Lisa", []string{"Donald"}),
	}

	l, r := 0, len(persons)-1
	for persons[l] != persons[r] {
		if persons[l].Knows(persons[r].name) {
			l++
		} else {
			r--
		}
	}

	fmt.Println("The celebrity is: ", persons[l].name)
}
