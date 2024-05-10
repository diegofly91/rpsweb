package rps

import (
	"math/rand"
	"strconv"
)

const (
	ROCK     = 0 // piedra vence a tijeras. (tijeras + 1) % 3 = 0
	PAPER    = 1 // papel vence a piedra. (piedra + 1) % 3 = 0
	SCISSORS = 2 // tijeras vencen a papel (papel + 1) % 3 = 0
)

type Round struct {
	Message           string `json:"message"`
	ComputerChoice    string `json:"computer_choice"`
	RountResult       string `json:"round_result"`
	ComputerChoiceInt int    `json:"computer_choice_int"`
	ComputerScore     string `json:"computer_score"`
	PlayerScore       string `json:"player_score"`
}

var winMessages = []string{
	"¡Ganaste!",
	"¡Buen Trabajo!",
	"Deberias Comprarte un boleto de loteria",
}

var loseMessages = []string{
	"¡Perdiste!",
	"¡Sigue Intentando!",
	"¡No te rindas!",
}

var drawMessages = []string{
	"Las grandes mentes, siempre piensan igual!",
	"¡Vuelve a intentarlo!",
	"¡No hay ganador!",
}

var ComputerScore, PlayerScore int

// Play is the function that determines the winner of the game
func PlayRound(player1 int) Round {
	computerValue := rand.Intn(3)

	var computerChoice, rountResult string
	var computerChoiceInt int

	switch computerValue {
	case ROCK:
		computerChoiceInt = ROCK
		computerChoice = "La computadora eligió Piedra"
	case PAPER:
		computerChoiceInt = PAPER
		computerChoice = "La computadora eligió Papel"
	case SCISSORS:
		computerChoiceInt = SCISSORS
		computerChoice = "La computadora eligió Tijeras"
	}

	messageInt := rand.Intn(3)
	var message string

	if player1 == computerValue {
		rountResult = "Es un empate"
		message = drawMessages[messageInt]
	} else if player1 == (computerValue+1)%3 {
		PlayerScore++
		rountResult = "¡El jugador a Ganado!"
		message = winMessages[messageInt]
	} else {
		ComputerScore++
		rountResult = "¡La computadora ha Ganado!"
		message = loseMessages[messageInt]
	}

	return Round{
		Message:           message,
		ComputerChoice:    computerChoice,
		RountResult:       rountResult,
		ComputerChoiceInt: computerChoiceInt,
		ComputerScore:     strconv.Itoa(ComputerScore),
		PlayerScore:       strconv.Itoa(PlayerScore),
	}
}
