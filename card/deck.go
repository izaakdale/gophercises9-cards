package deck

import (
	"math/rand"
	"time"
)

//go:generate go run golang.org/x/tools/cmd/stringer -type=SuitType
//go:generate go run golang.org/x/tools/cmd/stringer -type=RankType
type SuitType int
type RankType int

type Deck struct {
	Cards []Card
}

func NewDeck(shuffle bool, jacks int) Deck {
	var cards []Card
	// using less than hearts and king since these are the last iota
	// leaving jack for now
	for i := 1; i < int(Hearts)+1; i++ {
		for j := 1; j < int(King)+1; j++ {
			cards = append(cards, Card{
				Suit: SuitType(i),
				Rank: RankType(j),
			})
		}
		for i := 1; i < jacks; i++ {
			cards = append(cards, Card{
				Suit: SuitType(JokerS),
				Rank: RankType(JokerT),
			})
		}
	}

	ret := Deck{
		Cards: cards,
	}
	if shuffle != false {
		return ret.Shuffle()
	}
	return ret

}

func (d *Deck) Draw() Card {
	ret := d.Cards[0]
	d.Cards = append(d.Cards[1:], ret)
	return ret
}

func (d Deck) ShowDeck() {
	for _, c := range d.Cards {
		c.ShowCard()
	}
}

func (d Deck) Shuffle() Deck {

	var ret Deck

	rand.Seed(time.Now().Unix())
	rai := rand.Perm(len(d.Cards))
	for _, v := range rai {
		ret.Cards = append(ret.Cards, d.Cards[v])
	}
	return ret
}
