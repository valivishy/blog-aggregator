package handlers

import (
	"blog-aggregator/internal/commands"
	"context"
	"errors"
	"fmt"
)

func HandlerListFeeds(state *commands.State, _ commands.Command) error {
	feeds, err := state.DB.ListFeeds(context.Background())
	if err != nil {
		return errors.New("could not list feeds")
	}

	for _, feed := range feeds {
		fmt.Printf(" * %s(%s) created by %s\n", feed.FeedName, feed.FeedUrl, feed.UserName)
	}

	return nil
}
