package main

import "fmt"

type Person struct {
	name string
}

func (p Person) handleEvent(vacancices []string) {
	fmt.Println("Hey, ", p.name)
	fmt.Println("Here is something for you...")
	for _, vacancy := range vacancices {
		fmt.Println(vacancy)
	}
	fmt.Println()
}
