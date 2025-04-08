package handlers

import (
	"blog-aggregator/internal/commands"
	"context"
	"fmt"
)

func HandlerLogin(state *commands.State, command commands.Command) error {
	if len(command.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", command.Name)
	}
	name := command.Args[0]

	if _, err := state.DB.GetUser(context.Background(), name); err != nil {
		return fmt.Errorf("user %s doesn't exist", name)
	}

	err := state.Cfg.SetUser(name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Println("User switched successfully!")
	return nil
}
