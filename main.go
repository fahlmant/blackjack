package main

import (
	"bufio"
	"fmt"
	"math"
	"os"

	"github.com/fahlmant/blackjack/deck"
)

func main() {

	var d deck.Deck
	d.BuildDeck()
	d.Shuffle()

	totalMoney := 100

	for totalMoney >= 5 {

		// Ensure there's a full deck
		if len(d.Cards) < 17 {
			fmt.Println("Shuffling deck")
			d.BuildDeck()
			d.Shuffle()
		}

		fmt.Println("------------")

		// Set the bet
		fmt.Printf("Betting 5 dollars out of %d\n", totalMoney)
		bet := 5.0

		// Deal two cards to the player and dealer, and show them
		playerCards := []deck.Card{d.Draw()}
		dealerCards := []deck.Card{d.Draw()}
		playerCards = append(playerCards, d.Draw())
		dealerCards = append(dealerCards, d.Draw())
		printPlayerHand(playerCards)
		printDealerTopCard(dealerCards)

		pBJ := isBlackjack(playerCards)
		dBJ := isBlackjack(dealerCards)

		switch pBJ && dBJ {
		case true:
			fmt.Println("Push on 21")
			fmt.Println("Dealing new hand")
			fmt.Println()
			continue
		case false:
			if pBJ {
				// MFer taking the casino's hard earned dollars
				fmt.Println("Blackjack!!!")
				totalMoney += int(math.Floor(bet*.5)) + int(bet)
			} else if dBJ {
				// Suck it
				fmt.Println("Dealer has blackjack")
				totalMoney -= 5
			} else {
				break
			}

			fmt.Println("Dealing new hand")
			fmt.Println()
			continue
		}

		for {
			fmt.Printf("(h)it or (s)tand?\n>")
			var input string
			fmt.Scanf("%s", &input)
			if input == "h" {
				playerCards = append(playerCards, d.Draw())
				printPlayerHand(playerCards)
				if getHandTotal(playerCards) >= 21 {
					break
				}

				continue
			}
			if input == "s" {
				break
			}
		}

		playerTotal := getHandTotal(playerCards)

		if playerTotal > 21 {
			fmt.Println("Bust!")
			totalMoney -= int(bet)
			continue
		}
		fmt.Println("Player stands")
		printPlayerHand(playerCards)
		printDealerHand(dealerCards)

		for getHandTotal(dealerCards) < 17 {
			dealerCards = append(dealerCards, d.Draw())
			printDealerHand(dealerCards)
		}

		dealerTotal := getHandTotal(dealerCards)

		if dealerTotal > 21 {
			fmt.Println("Dealer busts")
			fmt.Println()
			totalMoney += int(bet)
			continue
		}

		if playerTotal > dealerTotal {
			fmt.Println("Player wins")
			fmt.Println()
			totalMoney += int(bet)
			continue
		}
		if playerTotal == dealerTotal {
			fmt.Println("Push")
			fmt.Println()
			continue
		}

		fmt.Println("Dealer wins")
		fmt.Println()

		// You get NOTHING! You LOSE! Good DAY sir!
		totalMoney -= int(bet)
		fmt.Printf("Press enter to deal new hand\n")
		input := bufio.NewScanner(os.Stdin)
		input.Scan()
	}

}

func printHand(name string, cards []deck.Card) {
	fmt.Printf("%s has: ", name)

	for _, card := range cards {
		fmt.Printf("%s of %s, ", card.FaceValue.Name, card.Suit)
	}
}

func printPlayerHand(cards []deck.Card) {

	printHand("Player", cards)
	fmt.Printf("for a total of %d\n", getHandTotal(cards))

}

func printDealerHand(cards []deck.Card) {
	printHand("Dealer", cards)
	fmt.Printf("for a total of %d\n", getHandTotal(cards))
}

func printDealerTopCard(cards []deck.Card) {
	fmt.Printf("Dealer has: %s of %s, for a total of %d\n", cards[0].FaceValue.Name, cards[0].Suit, cards[0].FaceValue.Value)
}

func getHandTotal(cards []deck.Card) int {
	localTotal := 0
	for _, card := range cards {
		localTotal += card.FaceValue.Value
	}

	return localTotal

}

func isBlackjack(cards []deck.Card) bool {

	// Extra check just to make sure that there are only 2 cards
	if len(cards) > 2 {
		return false
	}

	if getHandTotal(cards) == 21 {
		return true
	}

	return false
}
