package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	invalidUserInputError = "invalid user input"
	closingMessage        = "closing..."
)

func GetCoordinates() (bishopCoordinates, targetCoordinates []int, err error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter bishop coordinates or press 'e' to exit: ")
	bishop, err := reader.ReadString('\n')
	if err != nil {
		return nil, nil, err
	}

	// User wants to close the application
	if strings.TrimSpace(bishop) == "e" {
		return nil, nil, fmt.Errorf(closingMessage)
	}

	// Sanitize and validate the bishop coordinates
	bishopCleaned := strings.ReplaceAll(strings.TrimSpace(strings.ToLower(bishop)), " ", "")
	bishopCoordinates, err = ValidateAndTranlsateInputString(bishopCleaned)
	if err != nil {
		return nil, nil, err
	}

	fmt.Println("Enter target coordinates or press 'e' to exit: ")
	target, err := reader.ReadString('\n')
	if err != nil {
		return nil, nil, err
	}

	// User wants to close the application
	if strings.TrimSpace(target) == "e" {
		return nil, nil, fmt.Errorf(closingMessage)
	}

	// Sanitize and validate the target coordinates
	targetCleaned := strings.ReplaceAll(strings.TrimSpace(strings.ToLower(target)), " ", "")
	targetCoordinates, err = ValidateAndTranlsateInputString(targetCleaned)
	if err != nil {
		return nil, nil, err
	}

	// Final validation to prevent this edge case
	if bishopCoordinates[0] == targetCoordinates[0] && bishopCoordinates[1] == targetCoordinates[1] {
		return nil, nil, fmt.Errorf(invalidUserInputError)
	}

	fmt.Print("BISHOP: ", bishopCoordinates)
	fmt.Print("TARGET: ", targetCoordinates)

	return bishopCoordinates, targetCoordinates, nil
}

func ValidateAndTranlsateInputString(input string) ([]int, error) {
	// User input validation logic: no more than two characters
	if len(input) != 2 {
		return nil, fmt.Errorf(invalidUserInputError)
	}
	columnIndex := int(input[0] - 'a')
	rowIndex := int(input[1] - '1')

	// Validate coordinate for chess board size
	if columnIndex < 0 || columnIndex > 7 || rowIndex < 0 || rowIndex > 7 {
		return nil, fmt.Errorf(invalidUserInputError)
	}

	return []int{columnIndex, rowIndex}, nil
}
