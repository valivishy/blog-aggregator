package handlers

import (
	"blog-aggregator/internal/commands"
	"blog-aggregator/internal/database"
	"blog-aggregator/internal/rss"
	"context"
	"fmt"
	"github.com/google/uuid"
	"time"
)

const layout = "Mon, 02 Jan 2006 15:04:05 -0700"

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

	return scrapeFeed(feed, state)
}

func scrapeFeed(dbFeed database.Feed, state *commands.State) error {
	feed, err := rss.FetchFeed(context.Background(), dbFeed.Url)
	if err != nil {
		return err
	}

	fmt.Printf("Fetched feed %s:\n", feed.Channel.Title)

	for _, item := range feed.Channel.Item {
		if err := savePost(item, dbFeed.ID, state); err != nil {
			return err
		}
	}

	return nil
}

func savePost(item rss.Item, feedId uuid.UUID, state *commands.State) error {
	currentContext := context.Background()

	publishedAt, err := time.Parse(layout, item.PubDate)
	if err != nil {
		return err
	}

	_, err = state.DB.CreatePost(
		currentContext,
		database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       item.Title,
			Url:         item.Link,
			Description: item.Description,
			PublishedAt: publishedAt,
			FeedID:      feedId,
		},
	)

	return err
}
