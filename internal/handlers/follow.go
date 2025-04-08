package handlers

import (
	"blog-aggregator/internal/commands"
	"blog-aggregator/internal/database"
	"context"
	"fmt"
	"github.com/google/uuid"
	"time"
)

func HandlerFollow(state *commands.State, command commands.Command) error {
	args := command.Args
	if len(args) != 1 {
		return fmt.Errorf("expected 1 argument, got %d", len(args))
	}

	currentContext := context.Background()
	user, err := state.DB.GetUser(currentContext, state.Cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("user %s does not exist", state.Cfg.CurrentUserName)
	}

	url := args[0]
	feed, err := state.DB.GetFeedByUrl(currentContext, url)
	if err != nil {
		return err
	}

	follow, err := state.DB.CreateFeedFollow(
		currentContext,
		database.CreateFeedFollowParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			UserID:    user.ID,
			FeedID:    feed.ID,
		},
	)
	if err != nil {
		return err
	}

	fmt.Println(follow)

	return nil
}
