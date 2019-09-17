package main

import (
	"fmt"
)

type superHero struct {
	name  string
	power string
}

func new(name, power string) superHero {
	return superHero{name, power}
}

func (s superHero) action() {
	fmt.Printf("the super hero %s uses %s", s.name, s.action)
}

func (s superHero) mutation(name, power string) superHero {
	return new(name, power)
}

func main() {
	superman := new("Superman", "laser eyes")
	superman.action()
	supermanRage := superman.mutation("Superman - rage mode", "super laser eyes")
	superman.action()
	supermanRage.action()
}
