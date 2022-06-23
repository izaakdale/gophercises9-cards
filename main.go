package main

import (
	deck "github.com/izaakdale/deckOfCards/card"
)

func main() {

	d := deck.NewDeck(true, 2)
	d.Cards[3].ShowCard()
	// d.Cards[3].ShowCard()

	// sort.Sort(deck.ByRank(d.Cards))

	// d = d.Shuffle()
	// for _, c := range d.Cards {
	// 	deck.ShowCard(c)
	// }

}
