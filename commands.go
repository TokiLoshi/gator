package main

import (
	"fmt"
	// "github.com/TokiLoshi/gator/internal/config"
)


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



