package main

import (
	"fmt"
	"math/rand"
	"time"
)

type coin struct {
	sideUp string
}

type simulator interface {
	coinSide() string
}

func (c *coin) coinSide() string {
	rand.Seed(time.Now().UnixNano())

	min := 10
	max := 30

	randomNum := rand.Intn((max - min + 1) + min)
	if randomNum%2 == 0 {
		c.sideUp = "Heads"
	} else {
		c.sideUp = "Tails"
	}

	return c.sideUp
}

func info(s simulator) {
	fmt.Println("The coin landed on", s.coinSide())
}

func main() {
	c := coin{
		sideUp: "",
	}

	info(&c)
}
