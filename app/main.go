package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	for {
		fmt.Print("$ ")
		command, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}
		command = strings.TrimSpace(command)
		commands := strings.Fields(command)
		if commands[0] == "exit" {
			os.Exit(0)
		}
		if commands[0] == "echo" {
			if len(commands) > 1 {
				fmt.Printf("%s\n", strings.Join(commands[1:], " "))
			} else {
				fmt.Println("missing args.")
			}
		} else {
			fmt.Printf("%s: command not found\n", command)
		}
	}
}
