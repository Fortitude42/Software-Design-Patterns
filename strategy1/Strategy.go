package main

type Strategy interface {
	buildTask(word1 string, word2 string)
}
