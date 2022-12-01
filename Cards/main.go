package main

import "fmt"

func main() {
	cards := []string{"Ace of Diamonds", newCard()} //criando um Array
	cards = append(cards, "Six of Spades")          //incluindo um elemento no array

	fmt.Println(cards)

}

func newCard() string {
	return "Five of Diamonds"
}
