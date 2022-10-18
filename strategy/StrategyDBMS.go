package main

import "fmt"

type StrategyDBMS struct {
}

func (s StrategyDBMS) buildTask(word1 string, word2 string) {
	fmt.Println("____________________________")
	fmt.Println("DBMS task:")
	fmt.Println("What is the result of this code?")
	fmt.Println("SELECT * FROM " + word1)
	fmt.Println("RIGHT OUTER JOIN " + word2 + " ON...")
	fmt.Println("____________________________")
}
