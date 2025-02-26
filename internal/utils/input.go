package utils

import (
	"fmt"
	"strings"
	"syscall"

	"golang.org/x/term"
)

func GetSecureInput(prompt string) (string, error) {
	fmt.Print(prompt)
	defer fmt.Println()

	bytePassword, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return "", fmt.Errorf("failed to read cosmic secret: %w", err)
	}

	return string(bytePassword), nil
}

func Confirm(prompt string) bool {
	fmt.Printf("%s [y/N]: ", prompt)

	var response string
	fmt.Scanln(&response)

	return strings.ToLower(strings.TrimSpace(response)) == "y"
}
