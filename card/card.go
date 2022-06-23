package deck

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

type Card struct {
	Suit SuitType
	Rank RankType
}

func (c Card) ShowCard() {
	// fmt.Printf("%s of %s\n", c.Rank.String(), c.Suit.String())
	c.showCliCard()
}

func (c Card) GetScore() int {
	if int(c.Rank) > 1 && int(c.Rank) < 11 {
		// number cards
		return int(c.Rank)
	} else if int(c.Rank) > 10 {
		// picture cards score 10
		return 10
	} else if int(c.Rank) == 1 {
		// ace logic
		return 11
	} else {
		return 0
	}
}
