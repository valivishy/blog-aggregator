package handlers

import (
	"blog-aggregator/internal/commands"
	"blog-aggregator/internal/database"
	"context"
	"fmt"
	"github.com/google/uuid"
	"time"
)

func HandlerAddFeed(state *commands.State, command commands.Command) error {
	args := command.Args
	if len(args) != 2 {
		return fmt.Errorf("expected 2 arguments, got %d", len(args))
	}

	currentContext := context.Background()
	user, err := state.DB.GetUser(currentContext, state.Cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("user %s does not exist", state.Cfg.CurrentUserName)
	}

	feed, err := state.DB.CreateFeed(
		currentContext,
		database.CreateFeedParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Name:      args[0],
			Url:       args[1],
			UserID:    user.ID,
		},
	)
	if err != nil {
		return err
	}

	fmt.Println(feed)

	return nil
}
