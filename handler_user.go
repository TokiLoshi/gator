package main

import (
	"context"
	"fmt"
	"os"
)


func handlerLogin(s *state, cmd command) error {
	ctx := context.Background()
	queries := s.db
	username := cmd.Args[0]
	// If args slice is empyt return an error (should be a single argument, the username)

	if len(username) == 0 {
		return fmt.Errorf("command cannot be empty")
	} 

	user, err := queries.GetUser(ctx, username)
	if err != nil {
		fmt.Printf("user does not exists: %v and error: %v\n", user, err)
		os.Exit(1)
	}
	err = s.cfg.SetUser(username)
	if err != nil {
		return err
	}
	fmt.Println("User has been set")

	return nil
}