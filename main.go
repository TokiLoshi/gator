package main

import (
	"fmt"
	"os"

	"github.com/TokiLoshi/gator/internal/config"
)



func main() {

	cfg := &config.Config{}
	err := config.Read(cfg)
	if err != nil {
		fmt.Printf("Error reading the config")
		return
	}
	state := &State{configuration: cfg}

	commands := &Commands{
		handlers: map[string]func(*State, *Command) error{},
	}

	commands.register("login", handlerLogin)
	
	// cmd := &Command{name: "login", args: []string{"claireece"}}

	args := os.Args
	if len(args) <= 2 {
		fmt.Printf("%v arguments are too few, need at least 2\n", len(args))
		os.Exit(1)
	}
	fmt.Printf("user instructions: %v\n", args)
	argName := args[1]

	allArgs := args[2:]
	cmd := &Command{name: argName, args: allArgs}

	// Execute the command if it exists 
	if handler, exists := commands.handlers[cmd.name]; exists {
		err := handler(state, cmd)
		if err != nil {
			fmt.Printf("error executing command: %v\n", err)
			os.Exit(1)
		}
	} else {
		fmt.Println("Command not recognized.")
		os.Exit(1)
	}

	os.Exit(0)
}