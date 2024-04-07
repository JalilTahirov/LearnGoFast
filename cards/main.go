package main

import "fmt"

func main() {
	//var card string = "Ace of Spadec"
	card := newCard()
	fmt.Println(card)
}

func newCard() string {
	return "Five of Diamonds"
}
