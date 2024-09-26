package main

import (
	"fmt"
	// "github.com/TokiLoshi/gator/internal/config"
)

// type State struct {
// 	configuration *config.Config
// }

type command struct {
	Name string
	Args []string
}

// Stores the command name as keys and the handler function as a value
type commands struct {
	registeredCommands map[string]func(*state, command) error
}


func (c *commands) register(name string, fun func(*state, command) error) {
	c.registeredCommands[name] = fun
}

func (c *commands) run (s *state, cmd command) error {
	f, ok := c.registeredCommands[cmd.Name]
	if !ok {
		return fmt.Errorf("command not found")
	}
	return f(s, cmd)
}


// func handlerLogin(s *State, cmd *Command) error {
// 	fmt.Println("handleLogin")
// 	// If args slice is empyt return an error (should be a single argument, the username)
// 	if len(cmd.name) == 0 {
// 		return fmt.Errorf("command cannot be empty")
// 	} 
// 	command := cmd.args[0]
// 	err := s.configuration.SetUser(command)
// 	if err != nil {
// 		return err
// 	}
// 	fmt.Println("User has been set")

// 	return nil
// }

// func (c *Commands) register(name string, handler func(*State, *Command) error) {
// 	if c.handlers == nil {
// 		c.handlers = make(map[string]func(*State, *Command) error)
// 	}
// 	c.handlers[name] = handler

// }



