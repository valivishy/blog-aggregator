package handlers

import (
	"blog-aggregator/internal/commands"
	"blog-aggregator/internal/database"
	"context"
	"fmt"
)

func HandlerFollowing(state *commands.State, command commands.Command, user database.User) error {
	feeds, err := state.DB.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}

	fmt.Printf("%s is following:\n", user.Name)
	for _, feed := range feeds {
		fmt.Printf(" * %s\n", feed.FeedName)
	}

	return nil
}
