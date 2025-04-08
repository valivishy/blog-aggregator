package handlers

import (
	"blog-aggregator/internal/commands"
	"context"
	"fmt"
)

func HandlerFollowing(state *commands.State, command commands.Command) error {
	currentContext := context.Background()
	user, err := state.DB.GetUser(currentContext, state.Cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("user %s does not exist", state.Cfg.CurrentUserName)
	}

	feeds, err := state.DB.GetFeedFollowsForUser(currentContext, user.ID)
	if err != nil {
		return err
	}

	fmt.Printf("%s is following:\n", user.Name)
	for _, feed := range feeds {
		fmt.Printf(" * %s\n", feed.FeedName)
	}

	return nil
}
