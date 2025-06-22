package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/m4rkux/gator/internal/config"
)

type state struct {
	config config.Config
}

type command struct {
	name string
	args []string
}

var st *state

func main() {
	st = &state{}
	var err error
	(*st).config, err = config.Read()
	if err != nil {
		fmt.Println("Error reading the config file:", err)
	}

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := cleanInput(scanner.Text())

	switch input[0] {
	case "login":
		err = handlerLogin(st, command{
			name: "login",
			args: input,
		})

		if err != nil {
			fmt.Println(err)
		}
		break
	}
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("no username provided")
	}

	username := cmd.args[1]
	err := config.SetUser(username, s.config)
	if err != nil {
		fmt.Println("Error setting the current user:", err)
	}

	fmt.Printf("the %s user has been set.\n", username)
	return nil
}

func cleanInput(text string) []string {
	trimmed := strings.TrimSpace(text)
	lowered := strings.ToLower(trimmed)
	words := strings.Fields(lowered)
	return words
}
