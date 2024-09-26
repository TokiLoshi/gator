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

func getAllUsers(s *state, cmd command) error {
	fmt.Println("getting all users")
	if len(cmd.Args) > 2 {
		return fmt.Errorf("not enough arguments - must have at least 2")
	}
	ctx := context.Background()
	queries := s.db 
	users, err := queries.GetAllUsers(ctx)
	if err != nil {
		fmt.Println("error fetching users from queries")
		os.Exit(1)
	}

	for index, user := range users {
		if index == len(users) -1 {
			fmt.Println(user.Name + " (current)")
			continue
		}
		fmt.Println(user.Name)
	}
	
	os.Exit(0)
	return nil
}

func resetUserTable(s *state, cmd command) error {
	fmt.Println("restting all the users")
	if len(cmd.Args) > 2 {
		fmt.Printf("too many commands")
		os.Exit(1)
	}
	ctx := context.Background()
	queries := s.db
	err := queries.ResetUsers(ctx)
	if err != nil { 
		fmt.Printf("error resetting database %v", err)
		os.Exit(1)
	}
	
	os.Exit(0)
	return nil
}

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