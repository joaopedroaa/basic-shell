package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"strings"
	"time"
)

func main() {
	LineCharSymbol := "> "
	reader := bufio.NewReader(os.Stdin)

	for {
		getDirectory, _ := os.Getwd()
		getHostname, _ := os.Hostname()
		getUser, _ := user.Current()
		getTime := time.Now()

		var formatUser = getUser.Username + "@" + getHostname
		var formarHour = getTime.Format("15:04")
		var formatDirectory = strings.Split(getDirectory, "/")
		var formatDirectoryGetOneName = formatDirectory[len(formatDirectory)-1]

		fmt.Println(formarHour + " " + formatDirectoryGetOneName + " - " + formatUser + " ")
		fmt.Print(LineCharSymbol)

		input, err := reader.ReadString('\n')
		execComand := execInput(input, getUser)

		if err != nil {
			printErr(err)
		}

		if execComand != nil {
			printErr(execComand)
		}
	}
}

func printErr(err error) {
	fmt.Fprintln(os.Stderr, err)
}

func execInput(input string, user *user.User) error {
	input = strings.TrimSuffix(input, "\n")
	args := strings.Split(input, " ")

	switch args[0] {

	case "cd":
		if len(args) < 2 {
			return os.Chdir(user.HomeDir)
		}
		return os.Chdir(args[1])

	case "exit":
		os.Exit(0)

	}

	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
