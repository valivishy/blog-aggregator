package handlers

import (
	"blog-aggregator/internal/commands"
	"blog-aggregator/internal/rss"
	"context"
	"fmt"
	"time"
)

func HandlerAggregate(state *commands.State, command commands.Command) error {
	args := command.Args
	if len(args) != 1 {
		return fmt.Errorf("expected 1 argument, got %d", len(args))
	}

	timeBetweenRequests, err := time.ParseDuration(args[0])
	if err != nil {
		return err
	}

	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		if err := scrapeFeeds(state); err != nil {
			return err
		}
	}
}

func scrapeFeeds(state *commands.State) error {
	currentContext := context.Background()
	feed, err := state.DB.GetNextFeedToFetch(currentContext)
	if err != nil {
		return err
	}

	err = state.DB.MarkFeedFetched(currentContext, feed.ID)
	if err != nil {
		return err
	}

	return scrapeFeed(feed.Url)
}

func scrapeFeed(feedUrl string) error {
	feed, err := rss.FetchFeed(context.Background(), feedUrl)
	if err != nil {
		return err
	}

	fmt.Printf("Fetched feed %s:\n", feed.Channel.Title)
	for _, item := range feed.Channel.Item {
		fmt.Printf(" * %s:\n", item.Title)
	}

	return nil
}
