package main

import "fmt"

type StrategyDAA struct {
}

func (s StrategyDAA) buildTask(word1 string, word2 string) {
	fmt.Println("____________________________")
	fmt.Println("DAA task:")
	fmt.Println("What is next word?(Logic is easy actually):")
	fmt.Println(word1 + " --> " + word2 + "--> ...")
	fmt.Println("____________________________")
}
