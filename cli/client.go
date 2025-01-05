package cli

import (
	"fmt"
	"os"

	"github.com/los-dogos-studio/gurian-belote/game"
)

func StartBeloteGame(targetScore int) {
	fmt.Println("Starting Belote game")

	for {
		currentGame := game.NewBeloteGame(targetScore)
		playGame(&currentGame)
		if !askYesNo("Do you want to play again?") {
			break
		}
	}
}

func playGame(currentGame *game.BeloteGame) {
	currentGame.Start()
	for {
		switch currentGame.GetState() {
		case game.GameReady:
			panic("Invalid state reached")
		case game.GameInProgress:
			playStep(currentGame)
		case game.GameFinished:
			printGameResults(currentGame)
			return
		}
	}
}

func playStep(currentGame *game.BeloteGame) {
	handState := currentGame.GetHand().GetState()

	if handState != game.HandFinished {
		currentTurn, err := currentGame.GetHand().GetCurrentTurn()
		if err != nil {
			panic("Invalid error received when getting current turn:" + err.Error())
		}
		fmt.Printf("%v's turn\n", playerToString(currentTurn))
	}

	switch handState {
	case game.TableTrumpSelection:
		printGameResults(currentGame)
		playTableTrumpSelection(currentGame)
	case game.FreeTrumpSelection:
		playFreeTrumpSelection(currentGame)
	case game.HandInProgress:
		playHandInProgress(currentGame)
	}
}

func playTableTrumpSelection(currentGame *game.BeloteGame) {
	fmt.Println("Trump card on table: ", currentGame.GetHand().GetTableTrump())
	currentPlayer, err := currentGame.GetHand().GetCurrentTurn()
	if err != nil {
		panic("Invalid error received when getting current turn:" + err.Error())
	}

	printPlayerCards(currentGame.GetHand().GetPlayerCards(currentPlayer))

	res := askYesNo("Do you want to accept the trump card?")
	err = currentGame.AcceptTableTrump(currentPlayer, res)

	if err != nil {
		panic("Invalid error received when accepting table trump:" + err.Error())
	}
}

func playFreeTrumpSelection(currentGame *game.BeloteGame) {
	fmt.Println("Trump card on table: ", currentGame.GetHand().GetTableTrump())
	currentPlayer, err := currentGame.GetHand().GetCurrentTurn()
	if err != nil {
		panic("Invalid error received when getting current turn:" + err.Error())
	}

	printPlayerCards(currentGame.GetHand().GetPlayerCards(currentPlayer))

	res := askSkippableSuit("Do you want to choose trump suit?")
	err = currentGame.SelectTrump(currentPlayer, res)

	if err != nil {
		fmt.Fprintln(os.Stderr, "Please choose a valid suit")
	}
}

func playHandInProgress(currentGame *game.BeloteGame) {
	currentPlayer, err := currentGame.GetHand().GetCurrentTurn()
	if err != nil {
		panic("Invalid error received when getting current turn:" + err.Error())
	}

	fmt.Println("Trump suit:", currentGame.GetHand().GetTrump())
	printPlayerCards(currentGame.GetHand().GetPlayerCards(currentPlayer))
	printTableCards(currentGame.GetHand().GetTrick().GetTableCards())

	currentGame.GetHand().GetPlayerCards(currentPlayer)

	card := askCard("Play a card")
	err = currentGame.PlayCard(currentPlayer, card)

	if err != nil {
		fmt.Fprintln(os.Stderr, "Invalid card played,", err)
	}
}

func printGameResults(beloteGame *game.BeloteGame) {
	fmt.Println("Scores:")
	scores := beloteGame.GetScores()

	fmt.Println("Team 1: ", scores[game.Team1])
	fmt.Println("Team 2: ", scores[game.Team2])
}

func printPlayerCards(playerCards map[game.Card]bool) {
	fmt.Printf("Your cards: ")
	for card, owned := range playerCards {
		if owned {
			fmt.Printf("%s, ", card.String())
		}
	}
	fmt.Println()
}

func printTableCards(tableCards map[game.PlayerId]game.Card) {
	fmt.Println("Table cards: ")
	for player, card := range tableCards {
		fmt.Printf("--- %v: %s\n", playerToString(player), card.String())
	}
	fmt.Println()
}
