package handlers

import (
	"blog-aggregator/internal/commands"
	"blog-aggregator/internal/database"
	"context"
	"fmt"
	"github.com/google/uuid"
	"time"
)

func HandlerAddFeed(state *commands.State, command commands.Command, user database.User) error {
	args := command.Args
	if len(args) != 2 {
		return fmt.Errorf("expected 2 arguments, got %d", len(args))
	}

	feed, err := state.DB.CreateFeed(
		context.Background(),
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

	// We keep the second arg because we only need the URL
	command.Args = command.Args[1:]

	return HandlerFollow(state, command, user)
}
