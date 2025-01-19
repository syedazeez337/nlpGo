package main

import (
	"fmt"
	"strings"
)

// Lowercasing the given input
func toLower(s string) string {
	return strings.ToLower(s)
}

/*
func getInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter text (type 'exit' to quit):")

	for {
		fmt.Print("> ")
		if scanner.Scan() {
			input := strings.TrimSpace(scanner.Text())
			if input == "exit" {
				return fmt.Sprintln("Exiting...")
			}
			fmt.Println("You entered:", toLower(input))
			// return toLower(input)
		} else {
			// Handle error or EOF
			fmt.Println("Error reading input. Exiting...")
		}
	}
}
*/

func main() {
	text := "I am NeuralNets"
	text = toLower(text)
	fmt.Println(text)

}
