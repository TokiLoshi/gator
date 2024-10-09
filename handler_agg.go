package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/TokiLoshi/gator/internal/database"
	"github.com/google/uuid"
)

func handleBrowse(s *state, cmd command) error {
	if len(cmd.Args) < 1 || len(cmd.Args) > 2{
		fmt.Println("incorrect number of arguments")
		os.Exit(1)
	}
	limit64, err := strconv.ParseInt(cmd.Args[0], 10, 32)
	if err != nil {
		return fmt.Errorf("error parsing limit: %w",err)
	}
	limit := int32(limit64)
	fmt.Printf("Limit: %v\n", limit)
	ctx := context.Background() 
	queries := s.db 

	posts, err := queries.GetPosts(ctx, limit) 
	if err != nil {
		return fmt.Errorf("couldn't brows posts")
	}
	for index, post := range posts {
		fmt.Printf("%d - %v\n", index, post)
	}
	return nil
}

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
		pubDate, err := time.Parse(time.RFC1123Z, post.PubDate)
		if err != nil {
			fmt.Printf("error parsing publication date: %v\n", pubDate)
		}
		newPost, err := queries.CreatePost(ctx, database.CreatePostParams{
			ID: uuid.New(),
			CreatedAt: sql.NullTime{Time: time.Now(), Valid: true},
			UpdatedAt: sql.NullTime{Time: time.Now(), Valid: true},
			Title: sql.NullString{String: post.Title, Valid: post.Title != ""},
			Description: sql.NullString{String: post.Description, Valid: post.Description != ""},
			Url: sql.NullString{String: post.Link, Valid: post.Link != ""},
			PublishedAt: sql.NullTime{Time: pubDate, Valid: true},
			FeedID: feed.ID,
		})
		if err != nil {
			fmt.Printf("error saving new post to database: %v\n", err)
		}
		fmt.Printf("new post saved to database: %v\n", newPost.Title)
	}
}

