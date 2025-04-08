package handlers

import (
	"blog-aggregator/internal/commands"
	"context"
	"errors"
	"fmt"
)

func HandlerListUsers(state *commands.State, _ commands.Command) error {
	users, err := state.DB.ListUsers(context.Background())
	if err != nil {
		return errors.New("could not list users")
	}

	for _, user := range users {
		if user.Name == state.Cfg.CurrentUserName {
			fmt.Printf(" * %s (current)\n", user.Name)
		} else {
			fmt.Printf(" * %s\n", user.Name)
		}
	}

	return nil
}
