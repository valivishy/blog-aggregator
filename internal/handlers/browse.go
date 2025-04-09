package handlers

import (
	"blog-aggregator/internal/commands"
	"blog-aggregator/internal/database"
	"context"
	"fmt"
	"strconv"
)

func HandlerBrowse(state *commands.State, command commands.Command, user database.User) error {
	args := command.Args

	var limit int
	if len(args) != 1 {
		limit = 2
	} else {
		limit, _ = strconv.Atoi(args[0])
	}

	currentContext := context.Background()

	posts, err := state.DB.GetPostsForUser(
		currentContext,
		database.GetPostsForUserParams{
			UserID: user.ID,
			Limit:  int32(limit),
		},
	)

	if err != nil {
		return err
	}

	if len(posts) == 0 {
		fmt.Printf("%s does not yet have posts to browse\n", user.Name)
		return nil
	}

	fmt.Printf("%s is browsing:\n", user.Name)
	for _, post := range posts {
		fmt.Printf(" * %s\n", post.Title)
	}

	return nil
}
