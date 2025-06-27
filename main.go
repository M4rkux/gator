package main

import (
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

type commands map[string]func(*state, command) error

func (c *commands) run(s *state, cmd command) error {
	return (*c)[cmd.name](s, cmd)
}

func (c *commands) register(name string, f func(s *state, cmd command) error) {
	(*c)[name] = f
}

var st *state

func main() {
	st = &state{}
	var err error
	commands := commands{}

	commands.register("login", handlerLogin)

	(*st).config, err = config.Read()
	if err != nil {
		fmt.Println("Error reading the config file:", err)
	}

	if len(os.Args) <= 1 {
		fmt.Println("not enough arguments")
		os.Exit(1)
	}
	input := os.Args[1:]

	if _, ok := commands[input[0]]; ok {
		err = commands.run(st, command{
			name: input[0],
			args: input,
		})

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) <= 1 {
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

func cleanCommandInput(text string) string {
	trimmed := strings.TrimSpace(text)
	lowered := strings.ToLower(trimmed)
	return lowered
}
