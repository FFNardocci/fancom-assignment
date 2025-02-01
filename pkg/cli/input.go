package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetCoordinates() (bishop, target string, err error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter bishop coordinates: ")
	bishop, err = reader.ReadString('\n')
	if err != nil {
		return "", "", err
	}

	bishopValidated, err := ValidateInput(bishop)
	if err != nil {
		return "", "", err
	}
	bishopCleaned := strings.ReplaceAll(strings.TrimSpace(strings.ToLower(bishopValidated)), " ", "")

	fmt.Println("Enter target coordinates: ")
	target, err = reader.ReadString('\n')
	if err != nil {
		return "", "", err
	}

	targetValidated, err := ValidateInput(target)
	if err != nil {
		return "", "", err
	}
	targetCleaned := strings.ReplaceAll(strings.TrimSpace(strings.ToLower(targetValidated)), " ", "")

	return bishopCleaned, targetCleaned, nil
}

func ValidateInput(input string) (string, error) {
	// User input validation logic: no more than two characters, input must be within 8x8 board size

	return input, nil
}

func TranslateInput(input string) (coordinates []int) {
	columnIndex := int(input[0] - 'a')

	rowIndex := int(input[1] - '1')

	return []int{columnIndex, rowIndex}
}
