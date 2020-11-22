package main

import "fmt"

type bot interface {
	getGreetings() string
}

type englishBot struct{}
type urduBot struct{}

func main() {

	eb := englishBot{}
	ub := urduBot{}

	printGreetings(eb)
	printGreetings(ub)

}

func printGreetings(b bot) {
	fmt.Println(b.getGreetings())
}

func (englishBot) getGreetings() string {
	return "Hello !"
}

func (urduBot) getGreetings() string {
	return "Kia hal ha !"
}
