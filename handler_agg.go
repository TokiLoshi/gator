package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/TokiLoshi/gator/internal/database"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) < 1 || len(cmd.Args) >2 {
		fmt.Println("incorrect number of arguments")
		os.Exit(1)
	}
	args_time := cmd.Args[0]
	timeBetweenRequests, err := time.ParseDuration(args_time)
	if err != nil {
		return fmt.Errorf("error converting time: %w", err)
	}
	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}

func scrapeFeeds(s *state) {
	fmt.Println("scraping feeds")
	// get next feed from DB 
	ctx := context.Background()
	feed, err := s.db.GetNextFetched(ctx)
	if err != nil {
		fmt.Printf("couldn't get feed: %v", err)
		return
	}
	fmt.Printf("feed")
	scrapeFeed(s, feed)
} 

func scrapeFeed(s *state, feed database.Feed) {
	queries := s.db
	ctx := context.Background()
	feed, err := queries.GetNextFetched(ctx)
	if err != nil {
		fmt.Printf("error updating feed as fetched: %v", err)
		return
	}

	feedData, err := fetchFeed(ctx, feed.Url)
	if err != nil {
		fmt.Printf("error fetching feed %v", err)
	}
	for _, post := range feedData.Channel.Item {
		fmt.Printf("Post: %v\n", post.Title)
	}
}