package rps

import (
	"fmt"
	"testing"
)

func TestPlayRound(t *testing.T) {
	for i := 0; i < 3; i++ {
		round := PlayRound(i)

		switch i {
		case 0:
			fmt.Println("Testing PIEDRA")
		case 1:
			fmt.Println("Testing PAPEL")
		case 2:
			fmt.Println("Testing TIJERAS")
		}
		fmt.Printf("Message: %s\n", round.Message)
		fmt.Printf("ComputerChoice: %s\n", round.ComputerChoice)
		fmt.Printf("RoundResult: %s\n", round.RountResult)
		fmt.Printf("ComputerChoiceInt: %d\n", round.ComputerChoiceInt)
		fmt.Printf("ComputerScore: %s\n", round.ComputerScore)
		fmt.Printf("PlayerScore: %s\n", round.PlayerScore)
		fmt.Println("--------------------------------------------------")
	}
}
