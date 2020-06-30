package main

import (
	"fmt"
)

type animal interface {
	consumeFood()
	sleep()
	printStatus()
}

type bear struct {
	fishEaten  int32
	hoursSlept float32
}

type bull struct {
	grassKilosEaten int32
	hoursSlept      float32
}

// bears consume 5 fish
func (b *bear) consumeFood() {
	b.fishEaten += 5
}

// bulls consume 3 kilos of grass
func (b *bull) consumeFood() {
	b.grassKilosEaten += 3
}

// bears sleep a lot
func (b *bear) sleep() {
	b.hoursSlept += 10.4
}

// bulls sleep for a few hours
func (b *bull) sleep() {
	b.hoursSlept += 6.3
}

func (b bear) printStatus() {
	fmt.Printf("This bear has slept for a total of %f hours and eaten %d fish.\n", b.hoursSlept, b.fishEaten)
}

func (b bull) printStatus() {
	fmt.Printf("This bull has slept for a total of %f hours and eaten %d kilos of grass.\n", b.hoursSlept, b.grassKilosEaten)
}

// make the animal sleep twice and eat three times. Print the final status
func performActions(a animal) {
	a.consumeFood()
	a.consumeFood()
	a.sleep()
	a.consumeFood()
	a.sleep()
	a.printStatus()
}

func main() {
	bear1 := bear{hoursSlept: 0, fishEaten: 0}
	bull1 := bull{hoursSlept: 0, grassKilosEaten: 0}

	performActions(&bear1)
	performActions(&bull1)
}
