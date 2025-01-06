package cli

import (
	"fmt"
	"strings"

	"github.com/los-dogos-studio/gurian-belote/game"
)

func askYesNo(question string) bool {
	for {
		fmt.Print(question, " (y/n): ")
		var response string
		_, err := fmt.Scanln(&response)
		if err != nil {
			fmt.Println("Invalid input")
			continue
		}
		if response == "y" {
			return true
		} else if response == "n" {
			return false
		}
	}
}

func askSkippableSuit(question string) *game.Suit {
	for {
		fmt.Print(question, " (H/D/C/S/skip): ")
		var response string
		_, err := fmt.Scanln(&response)
		if err != nil {
			fmt.Println("Invalid input")
			continue
		}
		if response == "skip" {
			return nil
		}
		suit, err := stringToSuit(response)
		if err != nil {
			fmt.Println("Invalid suit")
			continue
		}
		return &suit
	}
}

func askCard(question string) game.Card {
	for {
		fmt.Print(question, " (e.g. 7H, 10D): ")
		var response string
		_, err := fmt.Scanln(&response)
		if err != nil {
			fmt.Println("Invalid input")
			continue
		}
		card, err := stringToCard(response)
		if err != nil {
			fmt.Println("Invalid card")
			continue
		}
		return card
	}
}

func stringToCard(s string) (game.Card, error) {
	rank, err := stringToRank(s[0 : len(s)-1])
	if err != nil {
		return game.Card{}, err
	}

	suit, err := stringToSuit(s[len(s)-1:])
	return game.Card{Rank: rank, Suit: suit}, err
}

func stringToRank(s string) (game.Rank, error) {
	switch strings.ToUpper(s) {
	case "7":
		return game.Seven, nil
	case "8":
		return game.Eight, nil
	case "9":
		return game.Nine, nil
	case "10":
		return game.Ten, nil
	case "J":
		return game.Jack, nil
	case "Q":
		return game.Queen, nil
	case "K":
		return game.King, nil
	case "A":
		return game.Ace, nil
	default:
		return game.Seven, fmt.Errorf("Invalid rank")
	}
}

func stringToSuit(s string) (game.Suit, error) {
	switch strings.ToUpper(s) {
	case "H":
		return game.Hearts, nil
	case "D":
		return game.Diamonds, nil
	case "C":
		return game.Clubs, nil
	case "S":
		return game.Spades, nil
	default:
		return game.Hearts, fmt.Errorf("Invalid suit")
	}
}

func playerToString(player game.PlayerId) string {
	switch player {
	case game.Player1:
		return "Player 1"
	case game.Player2:
		return "Player 2"
	case game.Player3:
		return "Player 3"
	case game.Player4:
		return "Player 4"
	default:
		panic("Invalid player id")
	}
}
