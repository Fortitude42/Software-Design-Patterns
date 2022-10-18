package main

import "fmt"

type StrategySoftwarePatterns struct {
}

func (s StrategySoftwarePatterns) buildTask(word1 string, word2 string) {
	fmt.Println("____________________________")
	fmt.Println("Software Patterns task:")
	fmt.Println("Implement this architecture in Golang:")
	fmt.Println("class " + word1 + " implements " + word2 + " {")
	fmt.Println("...")
	fmt.Println("}")
	fmt.Println("____________________________")
}
