package main

import (
	"fmt"

	"github.com/TokiLoshi/gator/internal/config"
)

type State struct {
	configuration *config.Config
}

type Command struct {
	name string
	args []string
}

// Stores the command name as keys and the handler function as a value
type Commands struct {
	handlers map[string]func(*State, *Command) error
}

func handlerLogin(s *State, cmd *Command) error {
	fmt.Println("handleLogin")
	// If args slice is empyt return an error (should be a single argument, the username)
	if len(cmd.name) == 0 {
		return fmt.Errorf("command cannot be empty")
	} 
	command := cmd.args[0]
	err := s.configuration.SetUser(command)
	if err != nil {
		return err
	}
	fmt.Println("User has been set")

	return nil
}

func (c *Commands) register(name string, handler func(*State, *Command) error) {
	if c.handlers == nil {
		c.handlers = make(map[string]func(*State, *Command) error)
	}
	c.handlers[name] = handler

}

