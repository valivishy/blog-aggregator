package handlers

import (
	"blog-aggregator/internal/commands"
	"blog-aggregator/internal/rss"
	"context"
	"fmt"
)

func HandlerAggregate(_ *commands.State, _ commands.Command) error {

	feed, err := rss.FetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return err
	}

	fmt.Println(feed)

	return nil
}
