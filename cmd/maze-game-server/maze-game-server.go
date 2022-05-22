package main

import (
	"flag"
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

	// Flags
	local := flag.Bool("local", false, "boolean")
	flag.Parse()

	// Create Logger
	l := log.New()

	// Create Config
	c := config.New(version, commit, date, *local)

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
