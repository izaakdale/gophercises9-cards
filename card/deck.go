package deck

import (
	"math/rand"
	"time"
)

//go:generate go run golang.org/x/tools/cmd/stringer -type=SuitType
//go:generate go run golang.org/x/tools/cmd/stringer -type=RankType
type SuitType int
type RankType int

const (
	JokerS SuitType = iota
	Spades
	Diamonds
	Clubs
	Hearts
)
const (
	JokerT RankType = iota
	Ace
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

type Deck struct {
	Cards []Card
}

type Card struct {
	Suit SuitType
	Rank RankType
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

func (c Card) ShowCard() {
	// fmt.Printf("%s of %s\n", c.Rank.String(), c.Suit.String())
	c.showCliCard()
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
