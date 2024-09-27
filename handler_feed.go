package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
	"os"
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
	fmt.Printf("URL: %v\n", url)

	ctx := context.Background()
	feed, err := fetchFeed(ctx, url)
	if err != nil {
		fmt.Printf("error fetching feed: %v\n", err)
	}
	fmt.Printf("feed: %v\n", feed)
	fmt.Printf("No errors, %v and feed: %v\n", err, feed)
	fmt.Println(feed.Channel.Title) 
	fmt.Println(feed.Channel.Description)
	fmt.Println(feed.Channel.Link)
	fmt.Printf("Number of items: %v\n", len(feed.Channel.Item))
	if len(feed.Channel.Item) > 0 {
		for i := range feed.Channel.Item {
			fmt.Printf(feed.Channel.Item[i].Title)
			fmt.Printf(feed.Channel.Item[i].Description)
		}
	}
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
	fmt.Printf("Response body first 20 chars: %s\n", string(body[:min(200, len(body))]))
	if err != nil {
		return nil, fmt.Errorf("error %v", err)
	}

	var feed RSSFeed 
	err = xml.Unmarshal(body, &feed)
	fmt.Printf("feed returned from Reading unmarshalled: %v\n", feed)
	if err != nil {
		fmt.Printf("error unmarshalling: %v\n", err)
		return nil, fmt.Errorf("error %v", err)
	}

	feed.Channel.Title = html.UnescapeString(feed.Channel.Title)
	fmt.Printf("Feed.title: %v\n", feed)
	feed.Channel.Description = html.UnescapeString(feed.Channel.Description)
	for i := range feed.Channel.Item {
		feed.Channel.Item[i].Title = html.UnescapeString(feed.Channel.Item[i].Title)
		feed.Channel.Item[i].Description = html.UnescapeString(feed.Channel.Item[i].Description)
	}
	return &feed, nil
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}