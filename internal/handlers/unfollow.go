package handlers

import (
	"blog-aggregator/internal/commands"
	"blog-aggregator/internal/database"
	"context"
	"fmt"
)

func HandlerUnfollow(state *commands.State, command commands.Command, user database.User) error {
	args := command.Args
	if len(args) != 1 {
		return fmt.Errorf("expected 1 argument, got %d", len(args))
	}

	currentContext := context.Background()
	url := args[0]

	return state.DB.DeleteFeedFollow(
		currentContext,
		database.DeleteFeedFollowParams{
			UserID: user.ID,
			Url:    url,
		},
	)
}
