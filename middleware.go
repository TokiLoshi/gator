package main

import (
	"context"
	"fmt"

	"github.com/TokiLoshi/gator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		ctx := context.Background()
		queries := s.db 
		currentUser := s.cfg.CurrentUserName 
		user, err := queries.GetUser(ctx, currentUser)
		if err != nil {
			return fmt.Errorf("error getting user")
		}
		return handler(s, cmd, user)
	}
}