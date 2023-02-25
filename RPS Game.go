package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	var playerMove string
	playerScore, compScore := 0, 0

	moves() // Shows the choices moves to input.

	// First one to get a score of 3 will be the final winner.
	for playerScore != 3 && compScore != 3 {
		compMove := computerMove() // Holds the return value from the computerMove function.

		// Shows the score of the player and computer.
		// strconv.Itoa() convert an integer into string to enable concatenation.
		fmt.Println("Player Score: " + strconv.Itoa(playerScore))
		fmt.Println("Computer Score: " + strconv.Itoa(compScore))
		fmt.Println()

		// Get player input move.
		fmt.Print("Enter your move: ")
		fmt.Scanln(&playerMove)
		fmt.Println()

		// Shows the move selected by the player and computer.
		fmt.Println("Player choose " + playerMove)
		fmt.Println("computer choose " + compMove)

		switch {
		case playerMove == compMove:
			fmt.Println("Its a draw")

		case playerMove == "r" && compMove == "s":
			fmt.Println("You win")
			playerScore++

		case playerMove == "p" && compMove == "r":
			fmt.Println("You win")
			playerScore++

		case playerMove == "s" && compMove == "p":
			fmt.Println("You win")
			playerScore++

		default:
			fmt.Println("Computer wins")
			compScore++
		}
		fmt.Println()
	}
}

func computerMove() string {
	var computerMove string
	seed := rand.NewSource(time.Now().UnixNano())
	randNum := rand.New(seed)

	randNumGen := randNum.Intn(4-1) + 1 // Random number generator range from 1 to 3.

	// For computer's random selection of moves.
	if randNumGen == 1 {
		computerMove = "r"
	} else if randNumGen == 2 {
		computerMove = "p"
	} else if randNumGen == 3 {
		computerMove = "s"
	}

	return computerMove
}

func moves() {
	fmt.Println("R O C K  P A P E R  S C I S S O R S  G A M E")
	fmt.Println("[r] - R O C K")
	fmt.Println("[p] - P A P E R")
	fmt.Println("[s] - S C I S S O R S")
	fmt.Println()
}
