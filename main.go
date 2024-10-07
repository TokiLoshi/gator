package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/TokiLoshi/gator/internal/config"
	"github.com/TokiLoshi/gator/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	db *database.Queries
	cfg *config.Config
}


func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("Error reading the config")
		os.Exit(1)
	}

	db, err := sql.Open("postgres", cfg.DBUrl)
	if err != nil {
		fmt.Println("Error opening database")
		os.Exit(1)
	}
	defer db.Close()
	dbQueries := database.New(db)

	programState := &state{
		cfg: &cfg, 
		db: dbQueries,
	
	}

	commands := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}

	commands.register("login", handlerLogin)
	commands.register("register", registerUser)
	commands.register("reset", resetUserTable)
	commands.register("users", getAllUsers)
	commands.register("agg", handleFeed)
		commands.register("feeds", getAllFeeds)
	// Needs middleware
	commands.register("addfeed", middlewareLoggedIn(handleAddFeed))
	commands.register("follow", middlewareLoggedIn(handleFollow))
	commands.register("unfollow", middlewareLoggedIn(handleUnfollow))
	commands.register("following", middlewareLoggedIn(getFollowing))

	// cmd := &Command{name: "login", args: []string{"claireece"}}

	args := os.Args
	// if len(args) <= 2 && (args[1] != "reset" || args[1] != "users") {
	// 	fmt.Printf("%v arguments are too few, need at least 2\n", len(args))
	// 	os.Exit(1)
	// }
	
	argName := args[1]
	allArgs := args[2:]

	err = commands.run(programState, command{Name: argName, Args: allArgs})
	if err != nil {
		fmt.Printf("error running command: %v\n", err)
		os.Exit(1)
	}
	os.Exit(0)
}