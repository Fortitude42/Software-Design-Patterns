package main

import "fmt"

type StrategyCalulus struct {
}

func (s StrategyCalulus) buildTask(word1 string, word2 string) {
	fmt.Println("____________________________")
	fmt.Println("Calculus 2 task:")
	fmt.Println("Calculate this homogeus function of 15th order:")
	fmt.Println(word1 + "(x)'' + " + word2 + "(x)''' = x^16 + cos(8x)")
	fmt.Println("____________________________")
}
