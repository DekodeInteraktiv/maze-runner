package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"time"

	"github.com/PeterBooker/maze-game-server/internal/config"
	"github.com/PeterBooker/maze-game-server/internal/log"
	"github.com/PeterBooker/maze-game-server/internal/server"
)

var (
	version string
	commit  string
	date    string
)

func main() {
	fmt.Println("Starting Maze Game")

	rand.Seed(time.Now().UnixNano())

	// Create Logger
	l := log.New()

	// Create Config
	c := config.New(version, commit, date)

	// Create Server
	s := server.New(l, c)

	// Graceful Shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	// Setup HTTP server.
	s.Setup()

	<-stop

	fmt.Println("Shutting Down Maze Game")
}
