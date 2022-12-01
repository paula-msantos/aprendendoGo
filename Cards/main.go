package main

import "fmt"

func main() {
	cards := deck{"Ace of Diamonds", newCard()} //criando um Array
	cards = append(cards, "Six of Spades")      //incluindo um elemento no array

	for i, card := range cards { // para iterar sobre um array fechado
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "Five of Diamonds"
}
