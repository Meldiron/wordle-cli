// Golang program to show how
// to take input from the user
package main

import (
	"fmt"
	"strings"
)

// main function
func main() {
	fmt.Print("Enter game-word: ")

	var word string
	fmt.Scanln(&word)

	for {
		fmt.Print("Enter guess: ")
		var guess string
		fmt.Scanln(&guess)

		for i, letter := range guess {
			currentLetter := string(letter)

			isGreen := false
			if i >= len(word) {
				isGreen = false
			} else {
				correctLetter := string(word[i])
				isGreen = currentLetter == correctLetter
			}

			isOrange := strings.Contains(word, currentLetter)

			if isGreen {
				fmt.Print("🟩")
			} else if isOrange {
				fmt.Print("🟧")
			} else {
				fmt.Print("⬛")
			}

		}

		fmt.Println("")
	}
}
