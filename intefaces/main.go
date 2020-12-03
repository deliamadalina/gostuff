package main

import "fmt"

type bot interface {
	getGreeding() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

}

func printGreeting(b bot) {
	fmt.Println(b.getGreeding())
}

func (englishBot) getGreeding() string {
	// very custom login for generating an english greeting
	return "Hi there!"
}

func (spanishBot) getGreeding() string {
	return "Hola!"
}
