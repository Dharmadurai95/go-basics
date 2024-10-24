package main

func main() {

	newDeck := newCards()

	newDeck.shuffle()
	newDeck.print()
	// newDeck.saveFile("_newDeck")
	// readFile("_newDeck")
}
