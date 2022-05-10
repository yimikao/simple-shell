package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

//keyboard is the standard input device(os.Stdin)

func main() {

	cmd_reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")

		input, err := cmd_reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		if err := execCmd(input); err != nil {
			fmt.Fprint(os.Stderr, err)
		}

	}

}

func execCmd(cmd string) error {

	cmd = strings.TrimSuffix(cmd, "\n")

	command := exec.Command(cmd)

	command.Stdin, command.Stderr = os.Stdin, os.Stderr

	return command.Run()

}
