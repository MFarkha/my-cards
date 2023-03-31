package main

// import "fmt"

func main() {
	// cards := newDeck()
	// cards.print()
	// hand, remainingDeck := deal(cards, 5)
	// hand.print()
	// remainingDeck.print()
	// fmt.Println(cards.toString())
	// cards.saveToFile("my_cards")
	new_cards := newDeckFromFile("my_cards")
	//	new_cards.print()
	new_cards.shuffle()
	new_cards.print()
}
