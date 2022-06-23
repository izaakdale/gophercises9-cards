package deck

import "fmt"

// 8 of spades should look like this...
//  __________
// | 8        |
// |          |
// |     S    |
// |    S     |
// |          |
// |        8 |
//  ----------

var topNumberStringJoker = "| :)       |"
var bottomNumberStringJoker = "|       :( |"
var topNumberStringA = "| Ace      |"
var bottomNumberStringA = "|        A |"
var topNumberString2 = "| 2        |"
var bottomNumberString2 = "|        2 |"
var topNumberString3 = "| 3        |"
var bottomNumberString3 = "|        3 |"
var topNumberString4 = "| 4        |"
var bottomNumberString4 = "|        4 |"
var topNumberString5 = "| 5        |"
var bottomNumberString5 = "|        5 |"
var topNumberString6 = "| 6        |"
var bottomNumberString6 = "|        6 |"
var topNumberString7 = "| 7        |"
var bottomNumberString7 = "|        7 |"
var topNumberString8 = "| 8        |"
var bottomNumberString8 = "|        8 |"
var topNumberString9 = "| 9        |"
var bottomNumberString9 = "|        9 |"
var topNumberString10 = "| 10       |"
var bottomNumberString10 = "|       10 |"
var topNumberStringJ = "| Jack     |"
var bottomNumberStringJ = "|        J |"
var topNumberStringQ = "| Queen    |"
var bottomNumberStringQ = "|        Q |"
var topNumberStringK = "| King     |"
var bottomNumberStringK = "|        K |"

var topSuitStringJoker = "|    jo    |"
var bottomSuitStringJoker = "|    ke    |"
var topSuitStringS = "|    Spa   |"
var bottomSuitStringS = "|    de    |"
var topSuitStringC = "|    Cl    |"
var bottomSuitStringC = "|    ub    |"
var topSuitStringD = "|    Dia   |"
var bottomSuitStringD = "|    mnd   |"
var topSuitStringH = "|    He    |"
var bottomSuitStringH = "|    art   |"

var paddingString = "|          |"
var topCardString = " __________"
var bottonCardString = " ----------"

func (c Card) getCliCardStrings() (string, string, string, string) {
	var topNumber, bottomNumber, topSuit, bottomSuit string

	switch c.Suit {
	case Spades:
		{
			topSuit = topSuitStringS
			bottomSuit = bottomSuitStringS
		}
	case Clubs:
		{
			topSuit = topSuitStringC
			bottomSuit = bottomSuitStringC
		}
	case Diamonds:
		{
			topSuit = topSuitStringD
			bottomSuit = bottomSuitStringD
		}
	case Hearts:
		{
			topSuit = topSuitStringH
			bottomSuit = bottomSuitStringH
		}
	case JokerS:
		{
			topSuit = topSuitStringJoker
			bottomSuit = bottomSuitStringJoker
		}
	}

	switch c.Rank {
	case Ace:
		{
			topNumber = topNumberStringA
			bottomNumber = bottomNumberStringA
		}
	case Two:
		{
			topNumber = topNumberString2
			bottomNumber = bottomNumberString2
		}
	case Three:
		{
			topNumber = topNumberString3
			bottomNumber = bottomNumberString3
		}
	case Four:
		{
			topNumber = topNumberString4
			bottomNumber = bottomNumberString4
		}
	case Five:
		{
			topNumber = topNumberString5
			bottomNumber = bottomNumberString5
		}
	case Six:
		{
			topNumber = topNumberString6
			bottomNumber = bottomNumberString6
		}
	case Seven:
		{
			topNumber = topNumberString7
			bottomNumber = bottomNumberString7
		}
	case Eight:
		{
			topNumber = topNumberString8
			bottomNumber = bottomNumberString8
		}
	case Nine:
		{
			topNumber = topNumberString9
			bottomNumber = bottomNumberString9
		}
	case Ten:
		{
			topNumber = topNumberString10
			bottomNumber = bottomNumberString10
		}
	case Jack:
		{
			topNumber = topNumberStringJ
			bottomNumber = bottomNumberStringJ
		}
	case Queen:
		{
			topNumber = topNumberStringQ
			bottomNumber = bottomNumberStringQ
		}
	case King:
		{
			topNumber = topNumberStringK
			bottomNumber = bottomNumberStringK
		}
	case JokerT:
		{
			topNumber = topNumberStringJoker
			bottomNumber = bottomNumberStringJoker
		}
	}

	return topNumber, bottomNumber, topSuit, bottomSuit
}

func (c Card) showCliCard() {
	topN, botN, topS, botS := c.getCliCardStrings()
	fmt.Printf("%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n",
		topCardString,
		topN,
		paddingString,
		topS,
		botS,
		paddingString,
		botN,
		bottonCardString,
	)
}
