package main

import (
	"cmd/main.go/pkg/cli"
	"cmd/main.go/pkg/logic"
	"fmt"
	"os"
)

func main() {
	for {
		if err := run(); err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			if err.Error() == "closing..." {
				break
			}
			os.Exit(1)
		}
	}
}

func run() error {
	// Gather and evaluate user input
	bishopCoordinates, targetCoordinates, err := cli.GetCoordinates()
	if err != nil {
		return err
	}

	// Evaluate if the coordinates are on the same diagonal
	if logic.CanBishopAttack(bishopCoordinates, targetCoordinates) {
		fmt.Println("They can attack")
	} else {
		fmt.Println("They can't attack")
	}

	return nil
}
