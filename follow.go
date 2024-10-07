package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/TokiLoshi/gator/internal/database"
	"github.com/google/uuid"
)

func handleFollow(s *state, cmd command, currentUser database.User) error {
	fmt.Printf("adding new follow")
	if len(cmd.Args) != 1 {
		return fmt.Errorf("not enough commands")
	}
	
	queries := s.db 
	ctx := context.Background()

	// current user 
	user := currentUser

	// feed to follow 
	url := cmd.Args[0] 
		
	fmt.Printf("Url: %v\n", url)
	
	// get url id: 
	feed_id, err := queries.GetFeedByUrl(ctx, url)
	if err != nil {
		return fmt.Errorf("error returning url id: %w", err)
	}

	// Create the record 
	newFollow, err := queries.CreateFeedFollow(ctx, database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: sql.NullTime{Time: time.Now(), Valid: true},
		UpdatedAt: sql.NullTime{Time: time.Now(), Valid: true},
		UserID: user.ID, 
		FeedID: feed_id.ID,
	})
	if err != nil {
		return fmt.Errorf("issue creating newFollow %v : %w", newFollow, err)
	}
	// create a new feed follow record for current user
	fmt.Printf("new follow record created: %v\n", newFollow)
	feedName := feed_id.Name
	userName := user.Name
	// Print the name of the feed
	fmt.Printf("%v is now following the feed %v\n", userName, feedName)
	
	return nil
}

func getFollowing(s *state, cmd command, currentUser database.User) error {

	ctx := context.Background()
	queries := s.db 

	user := currentUser
	user_id := user.ID 
	follows, err := queries.GetFeedFollowsForUser(ctx, user_id)
	if err != nil {
		return fmt.Errorf("error getting following for feed: %w", err)
	}
	for _, follow := range follows {
		fmt.Printf("%v\n", follow.FeedName)
	}
	return nil
}