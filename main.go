package main

func main() {
	//var card string = "Ace of Spades"

	//card := "Ace of Spades"
	//card = "Five of diamonds"

	//card := newCard()
	//
	//fmt.Println(card)

	//cards := deck{"Ace of Diamonds", newCard()}
	//cards = append(cards, "Six of Spades")

	//for i, card := range cards {
	//	fmt.Println(i, card)
	//}

	//cards := newDeck()
	//cards.print()

	//hand, remainingCards := deal(cards, 5)

	//hand.print()
	//remainingCards.print()

	// SAVE DECK FROM FILE
	//cards.saveToFile("cosik.txt")

	// READ DECK FROM FILE
	//newCards := newDeckFromFile("cosik.txt")
	//newCards.print()

	// SHUFFLE
	cards := newDeck()
	cards.shuffle()
	cards.print()

	//fmt.Println(hand)
	//fmt.Println(remainingCards)

	// BYTE SLICE
	//test := "hello there"
	//fmt.Println([]byte(test))
}

//func newCard() string {
//	return "Five of Diamonds"
//}
