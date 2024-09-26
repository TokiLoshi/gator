package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/TokiLoshi/gator/internal/database"
	"github.com/google/uuid"
)

func registerUser(s *state, cmd command) error {
	fmt.Println("tyring to register user")
	ctx := context.Background()
	queries := s.db
	username := cmd.Args[0]

	user, err := queries.GetUser(ctx, username)
	if err == nil {
		fmt.Printf("user already exists: %v and error: %v\n", user, err)
		os.Exit(1)
	}

	newUser, err := queries.CreateUser(ctx, database.CreateUserParams {
		ID : uuid.New(),
		CreatedAt: sql.NullTime{Time: time.Now(), Valid: true},
		UpdatedAt: sql.NullTime{Time: time.Now(), Valid: true},
		Name : username,
	}) 
	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("new user created: %v\n", newUser)
	err = handlerLogin(s, cmd)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("logged in: %v\n", username)



	return nil
}