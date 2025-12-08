package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Print("$ ")
	command, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		log.Fatal("something went wrong")
	}
	command = strings.TrimSpace(command)
	fmt.Printf("%s: command not found\n", command)
}
