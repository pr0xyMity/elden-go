package main

func main() {
	cards := deck{"Ace of Heart", newCard()}
    cards = append(cards, "Six of Spades")
    cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
