package handlers

import (
	"blog-aggregator/internal/commands"
	"context"
	"fmt"
)

func HandlerReset(state *commands.State, _ commands.Command) error {
	err := state.DB.DeleteUsers(context.Background())
	if err != nil {
		return err
	}

	fmt.Println("Users dropped successfully!")

	return nil
}
