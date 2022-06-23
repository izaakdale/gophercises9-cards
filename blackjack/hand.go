package blackjack

import (
	deck "github.com/izaakdale/deckOfCards/card"
)

type Hand struct {
	Score int
	Cards []deck.Card
}

func (h Hand) Blackjack() bool {
	if h.Score < 21 || h.Score > 21 {
		return false
	} else {
		// score is 21 check for blackjack
		if h.Cards[0].GetScore() >= 11 && h.Cards[1].GetScore() >= 10 {
			return true
		} else if h.Cards[1].GetScore() >= 11 && h.Cards[0].GetScore() >= 10 {
			return true
		} else {
			return false
		}
	}
}

func (h *Hand) AddToScore(c deck.Card) {
	// if card is ace score is over 10 then add one to score,
	// otherwise add the standard score
	if int(c.Rank) == 1 && h.Score > 10 {
		h.Score += 1
	} else {
		h.Score += c.GetScore()
	}
}

func (h *Hand) Draw(d *deck.Deck, show bool) {
	c := d.Draw()
	h.Cards = append(h.Cards, c)

	// add to score
	h.AddToScore(c)
	if show {
		c.ShowCard()
	}
}
