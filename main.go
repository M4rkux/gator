package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/m4rkux/gator/internal/config"
	"github.com/m4rkux/gator/internal/database"
)

type state struct {
	db  *database.Queries
	cfg config.Config
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
	commands.register("register", handlerRegister)
	commands.register("reset", handlerReset)
	commands.register("users", handlerListUsers)

	(*st).cfg, err = config.Read()
	if err != nil {
		fmt.Println("Error reading the config file:", err)
	}

	db, err := sql.Open("postgres", (*st).cfg.DbUrl)
	if err != nil {
		fmt.Println("Error connectin to database")
		os.Exit(1)
	}

	(*st).db = database.New(db)

	if len(os.Args) <= 1 {
		fmt.Println("not enough arguments")
		os.Exit(1)
	}
	input := os.Args[1:]

	cmdName := cleanCommandInput(input[0])

	if _, ok := commands[cmdName]; ok {
		err = commands.run(st, command{
			name: cmdName,
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

	_, err := (*s).db.GetUser(context.Background(), username)
	if err != nil {
		return errors.New("User not found")
	}

	err = config.SetUser(username, s.cfg)
	if err != nil {
		return err
	}

	fmt.Printf("the %s user has been set.\n", username)
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) <= 1 {
		return errors.New("no user provided")
	}

	username := cmd.args[1]

	_, err := (*s).db.GetUser(context.Background(), username)
	if err == nil {
		return errors.New("User already created")
	}

	user, err := (*s).db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      username,
	})
	if err != nil {
		return err
	}

	err = config.SetUser(username, s.cfg)
	if err != nil {
		return err
	}

	fmt.Println(user)

	return nil
}

func handlerReset(s *state, cmd command) error {
	fmt.Println("Reseting database..")
	err := (*s).db.DeleteAllUsers(context.Background())
	if err != nil {
		return err
	}

	fmt.Println("All users were deleted")
	return nil
}

func handlerListUsers(s *state, cmd command) error {

	users, err := (*s).db.GetUsers(context.Background())
	if err != nil {
		return err
	}

	for _, user := range users {
		item := fmt.Sprintf(" * %s", user.Name)

		if (*s).cfg.CurrentUserName == user.Name {
			item += " (current)"
		}

		fmt.Println(item)
	}

	return nil
}

func cleanCommandInput(text string) string {
	trimmed := strings.TrimSpace(text)
	lowered := strings.ToLower(trimmed)
	return lowered
}
