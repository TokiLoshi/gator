package main

import (
	"context"
	"database/sql"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/TokiLoshi/gator/internal/database"
	"github.com/google/uuid"
)

type RSSFeed struct {
	XMLName xml.Name `xml:"rss"`
	Channel struct {
		Title string `xml:"title"`
		Link string `xml:"link"`
		Description string `xml:"description"`
		Item []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title string `xml:"title"`
	Link string `xml:"link"`
	Description string `xml:"description"`
	PubDate string `xml:"pubDate"`
}

func handleFeed(s *state, cmd command) error {
	if len(cmd.Args) > 3 {
		fmt.Println("error in feed- too many arguments")
		os.Exit(1)
	}
	url := "https://www.wagslane.dev/index.xml"

	ctx := context.Background()
	feed, err := fetchFeed(ctx, url)
	if err != nil {
		fmt.Printf("error fetching feed: %v\n", err)
	}
	
	if len(feed.Channel.Item) > 0 {
	for i := range feed.Channel.Item {
		feed.Channel.Item[i].Title = html.UnescapeString(feed.Channel.Item[i].Title)
		feed.Channel.Item[i].Description = html.UnescapeString(feed.Channel.Item[i].Description)
		fmt.Println(feed.Channel.Item[i].Title)
		fmt.Println(feed.Channel.Item[i].Description)
	}
	}
	os.Exit(0)
	return nil
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	fmt.Println("fetching RSS Feeed")
	
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, feedURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error getting feed")
	}

	req.Header.Set("User-Agent", "gator")

	client := &http.Client{}
	
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error fetching from client")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error %v", err)
	}

	var feed RSSFeed 
	err = xml.Unmarshal(body, &feed)

	if err != nil {
		fmt.Printf("error unmarshalling: %v\n", err)
		return nil, fmt.Errorf("error %v", err)
	}

	return &feed, nil
}

func handleAddFeed(s *state, cmd command, currentUser database.User) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("not enough commands ")
	}
	name := cmd.Args[0]
	url := cmd.Args[1]
	user := currentUser
	ctx := context.Background()
	queries := s.db
	fmt.Printf("Current user: %v\n", user)
	newFeed, err := queries.CreateFeed(ctx, database.CreateFeedParams {
		ID: uuid.New(),
		CreatedAt: sql.NullTime{Time: time.Now(), Valid: true},
		UpdatedAt: sql.NullTime{Time: time.Now(), Valid: true},
		Name: name, 
		Url: url,
		UserID: user.ID,
	})
	if err != nil {
		return fmt.Errorf("error creating feed %w", err)
	}
	fmt.Printf("New feed: %v", newFeed)
	fmt.Printf("Automatically adding follow")

	followCmd := command {
		Name: "follow",
		Args: []string{url},
	}
	err = handleFollow(s, followCmd, user)
	if err != nil {
		return fmt.Errorf("issue auto following: %w", err)
	}
	return nil

}

// Takes no arguments, prints all feeds in db
// Include the name of the feed 
// Url of the feed 
// name of the user (might need a new sql query)

func getAllFeeds(s *state, cmd command) error {
	if len(cmd.Args) > 1 {
		return fmt.Errorf("not enough arguments - must have 2")
	}
	ctx := context.Background()
	queries := s.db
	feeds, err := queries.GetFeeds(ctx)
	if err != nil {
		return fmt.Errorf("error getting all feeds: %w", err)
	}

	for _, feed := range feeds {
		fmt.Printf("%v\n", feed.Name)
		fmt.Printf("%v\n", feed.Url)
		username, err := queries.GetUserById(ctx, feed.UserID)
		if err != nil {
			fmt.Print("couldn't find user here")
			continue
		}
		fmt.Printf("%v\n", username.Name)
	} 

	return nil
} 
