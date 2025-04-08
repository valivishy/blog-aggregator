package middleware

import (
	"blog-aggregator/internal/commands"
	"blog-aggregator/internal/database"
	"context"
	"fmt"
)

func LoggedIn(handler func(state *commands.State, cmd commands.Command, user database.User) error) func(*commands.State, commands.Command) error {
	return func(state *commands.State, cmd commands.Command) error {
		user, err := state.DB.GetUser(context.Background(), state.Cfg.CurrentUserName)
		if err != nil {
			return fmt.Errorf("user %s does not exist", state.Cfg.CurrentUserName)
		}

		return handler(state, cmd, user)
	}
}
