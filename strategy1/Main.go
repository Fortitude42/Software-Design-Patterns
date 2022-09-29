package main

import "fmt"

func main() {
	var word1, word2 string
	fmt.Println("Enter two words to build a task")
	fmt.Println("Enter your first word:")
	fmt.Scanln(&word1)

	fmt.Println("Enter your second word:")
	fmt.Scanln(&word2)

	fmt.Println()

	p := Professor{}
	p.SetStrategy(StrategySoftwarePatterns{})
	p.strategy.buildTask(word1, word2)

	p.SetStrategy(StrategyCalulus{})
	p.strategy.buildTask(word1, word2)

	p.SetStrategy(StrategyDBMS{})
	p.strategy.buildTask(word1, word2)

	p.SetStrategy(StrategyDAA{})
	p.strategy.buildTask(word1, word2)

}
