package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	cliCommands := getCliCommands()
	for {
		fmt.Print("Pokedex > ")
		ok := scanner.Scan()
		if !ok {
			break
		}

		command := scanner.Text()
		cli, ok := cliCommands[command]
		if !ok {
			continue
		}

		err := cli.callback()

		if err != nil {
			break
		}
	}
}
