package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	LineCharSymbol := "> "

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(LineCharSymbol)
		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println(err)
		}

		if err = execInput(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

	}
}

func execInput(input string) error {
	// Remove the newline character.
	input = strings.TrimSuffix(input, "\n")

	// Prepare the command to execute.
	cmd := exec.Command(input)
	// Set the correct output device.
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	// Execute the command and return the error.
	return cmd.Run()
}
