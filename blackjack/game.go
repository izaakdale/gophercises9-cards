package blackjack

import (
	"fmt"
	"strings"
	"time"

	deck "github.com/izaakdale/deckOfCards/card"
)

func StartGame() {

	d := deck.NewDeck(true, 0)

	for {
		dealer := Hand{}
		player := Hand{}

		fmt.Printf("Dealer Card\n")
		dealer.Draw(&d, true)
		time.Sleep(time.Second * 2)

		fmt.Printf("Player Cards\n")
		player.Draw(&d, true)
		time.Sleep(time.Second * 2)

		dealer.Draw(&d, false)

		player.Draw(&d, true)

		// now start the actual gambling
		for {
			if player.Score > 21 {
				fmt.Println("BUST!")
				break
			} else if player.Blackjack() {
				fmt.Println("You won with a blackjack!")
				break
			}

			fmt.Printf("Score: %d\n", player.Score)
			fmt.Print("Would you like to hit or stick: ")

			var choice string
			// Taking input from user
			fmt.Scanln(&choice)
			if strings.ToLower(choice) == "hit" || strings.ToLower(choice) == "h" {
				time.Sleep(time.Second / 2)
				player.Draw(&d, true)
				if dealer.Score < 17 {
					dealer.Draw(&d, false)
				}
			} else if strings.ToLower(choice) == "stick" || strings.ToLower(choice) == "s" {
				// winning logic
				if player.Score > dealer.Score && dealer.Score <= 21 {
					fmt.Printf("You won with score of %d against %d\n", player.Score, dealer.Score)
				} else if player.Score < dealer.Score && dealer.Score <= 21 {
					fmt.Printf("You lost with score of %d against %d\n", player.Score, dealer.Score)
				} else if player.Score == dealer.Score && dealer.Score <= 21 {
					fmt.Printf("Draw with score of %d against %d\n", player.Score, dealer.Score)
				} else {
					// only way it hits here is if dealer score is above 21
					fmt.Printf("You won with score of %d against %d\n", player.Score, dealer.Score)
				}
				break
			} else {
				fmt.Println("Please enter hit, h, stick or s")
			}
		}
		fmt.Print("Play again with same deck? Q to cancel: ")
		var answer string
		fmt.Scanln(&answer)
		if strings.ToLower(answer) == "q" {
			break
		}
	}
}
