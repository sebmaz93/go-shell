package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Command struct {
	Name        string
	Description string
	Handler     func(args, dirs []string) error
}

var builtInCommands map[string]Command

func init() {
	builtInCommands = map[string]Command{
		"exit": {
			Name:        "exit",
			Description: "Exit the app.",
			Handler:     exitHandler,
		},
		"echo": {
			Name:        "echo",
			Description: "Echo the command arguments.",
			Handler:     echoHandler,
		},
		"type": {
			Name:        "type",
			Description: "Describe a command.",
			Handler:     typeHandler,
		},
	}
}

func exitHandler(args, dirs []string) error {
	os.Exit(0)
	return nil
}

func echoHandler(args, dirs []string) error {
	if len(args) == 0 {
		return fmt.Errorf("missing args, command is: echo <args>\n")
	}
	fmt.Printf("%s\n", strings.Join(args, " "))
	return nil
}

func typeHandler(args, dirs []string) error {
	if len(args) == 0 {
		return fmt.Errorf("missing args, command is: type <args>")
	}

	cmd := args[0]

	if _, ok := builtInCommands[cmd]; ok {
		fmt.Printf("%s is a shell builtin\n", cmd)
		return nil
	}

	for _, dir := range dirs {
		fullPath := filepath.Join(dir, cmd)

		info, err := os.Stat(fullPath)
		if err != nil {
			continue
		}

		mode := info.Mode()

		if !mode.IsDir() && mode&0111 != 0 {
			fmt.Printf("%s is %s\n", cmd, fullPath)
			return nil
		}
	}

	fmt.Printf("%s: not found\n", cmd)
	return nil
}

func main() {
	paths := os.Getenv("PATH")
	dirs := filepath.SplitList(paths)

	for {
		fmt.Print("$ ")
		command, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}
		command = strings.TrimSpace(command)
		tokens := strings.Fields(command)
		name := tokens[0]
		args := tokens[1:]

		cmd, ok := builtInCommands[name]
		if !ok {
			fmt.Printf("%s: command not found\n", name)
			continue
		}

		if err := cmd.Handler(args, dirs); err != nil {
			fmt.Fprint(os.Stderr, "Error: ", err)
		}
	}
}
