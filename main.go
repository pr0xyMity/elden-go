package main

import "fmt"

func main() {
    cards := newDeck()
    cards.print()

    hand, remainingCards := deal(cards, 3)
    fmt.Println(hand, remainingCards)
}
