package main

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

	programState := &commands.State{
		Cfg: cfg,
	}

	localCommands := commands.Commands{
		RegisteredCommands: make(map[string]func(*commands.State, commands.Command) error),
	}
	localCommands.Register("login", handlers.HandlerLogin)

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
