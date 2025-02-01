package main

import (
	"cmd/main.go/pkg/cli"
	"cmd/main.go/pkg/logic"
	"fmt"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	// Initialize the board
	board := logic.InitBoard(8)
	fmt.Println(board)

	// Gather and evaluate user input
	bishop, target, err := cli.GetCoordinates()
	if err != nil {
		return err
	}
	fmt.Printf("Received chess notation: %v - %v\n", bishop, target)

	// Translate chess notation into numerical indices
	bishopCoordinates := cli.TranslateInput(bishop)
	fmt.Println("Bishop coordinates: ", bishopCoordinates)

	targetCoordinates := cli.TranslateInput(target)
	fmt.Println("Target coordinates: ", targetCoordinates)

	// Evaluate if the coordinates are on the same diagonal
	if logic.CanBishopAttack(bishopCoordinates, targetCoordinates) {
		fmt.Println("They can attack")
	} else {
		fmt.Println("They can't attack")
	}
	return nil
}
