package main

type Observable interface {
	subscribe(observer Observer)
	sendAll()
}
