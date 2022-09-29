package main

type Professor struct {
	strategy Strategy
}

func (p *Professor) SetStrategy(strategy Strategy) {
	p.strategy = strategy
}
