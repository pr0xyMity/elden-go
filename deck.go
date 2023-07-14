package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []card

// fmt.Printf("%+v", value) <= this %+v prints all the keys not only values
type card struct {
    suit string
    value string
}

func newDeck() deck {
    cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
    cardValues := []string{"Ace", "Two", "Three", "Four"}

    cards := deck{}

    for _, suit := range cardSuits {
        for _, value := range cardValues {
            singleCard := card{value, suit}
            cards = append(cards, singleCard)        
        }
    }

    return cards 
}

// This is a function RECEIVER 
// it means that any variable of type "deck"
// now has the access to the "print" method
func (d deck) print() {
    for i, card := range d {
        fmt.Println(i, card)
    }
}

func deal(d deck, handSize int) (deck, deck) {
    return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
    return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
    return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
    bs, err := ioutil.ReadFile(filename)
    if err != nil {
       os.Exit(1) 
    }

    s := strings.Split(string(bs), ",")
    return deck(s)
}

func (d deck) shuffle() {
    source := rand.NewSource(time.Now().UnixNano())
    r := rand.New(source)

    for i:= range d {

        newPosition := r.Intn(len(d) - 1) 

        d[i], d[newPosition] = d[newPosition], d[i]
    }
}
