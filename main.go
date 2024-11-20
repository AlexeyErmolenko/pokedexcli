package main

import (
	"bufio"
	"fmt"
	"os"
)

var escseq = "\033["
var tealColor, blueColor, greenColor, redColor = "36", "34", "32", "31"
var regularStyle = "0"
var end = "m"

var divider = ";"

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	cliCommands := getCliCommands()
	for {
		fmt.Printf("%sPokedex %s> %s", escseq+tealColor+end, escseq+blueColor+end, escseq+greenColor+end)
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
			if err.Error() == "exit status 1" {
				break
			} else {
				fmt.Printf("%s%w%s\n", escseq+redColor+end, err, escseq+regularStyle+end)
			}
		}
	}
}
