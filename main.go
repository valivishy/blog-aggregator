package main

import (
	"blog-aggregator/internal/database"
	"blog-aggregator/internal/middleware"
	"database/sql"
	_ "github.com/lib/pq"
)

import (
	"blog-aggregator/internal/commands"
	"blog-aggregator/internal/config"
	"blog-aggregator/internal/handlers"
	"log"
	"os"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	db, err := sql.Open("postgres", cfg.DbUrl)
	if err != nil {
		log.Fatalf("error connecting to the DB: %v", err)
	}

	programState := &commands.State{
		Cfg: cfg,
		DB:  database.New(db),
	}

	localCommands := commands.Commands{
		RegisteredCommands: make(map[string]func(*commands.State, commands.Command) error),
	}

	localCommands.Register("login", handlers.HandlerLogin)
	localCommands.Register("register", handlers.HandlerRegister)
	localCommands.Register("reset", handlers.HandlerReset)
	localCommands.Register("users", handlers.HandlerListUsers)
	localCommands.Register("agg", handlers.HandlerAggregate)
	localCommands.Register("feeds", handlers.HandlerListFeeds)

	localCommands.Register("addfeed", middleware.LoggedIn(handlers.HandlerAddFeed))
	localCommands.Register("follow", middleware.LoggedIn(handlers.HandlerFollow))
	localCommands.Register("following", middleware.LoggedIn(handlers.HandlerFollowing))
	localCommands.Register("unfollow", middleware.LoggedIn(handlers.HandlerUnfollow))

	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
		return
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	err = localCommands.Run(programState, commands.Command{Name: cmdName, Args: cmdArgs})
	if err != nil {
		log.Fatal(err)
	}
}
