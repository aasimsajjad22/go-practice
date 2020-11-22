package main

func main() {

	//var card string = "Ace of cards"
	//cards := []string{"Ace of Diamonds", newCard()}

	//cards := deck{"Ace of Diamonds", newCard()}

	cards := newDeck()

	//hand, remainingCards := deal(cards, 5)

	//hand.print()
	//remainingCards.print()
	//fmt.Println(cards.toString())
	//cards.saveToFile("mycards")
	//cards := newDeckFromFile("mycards")

	cards.shuffle()
	cards.print()
}
