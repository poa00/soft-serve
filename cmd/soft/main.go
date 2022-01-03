package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/charmbracelet/soft-serve/config"
	"github.com/charmbracelet/soft-serve/server"
)

var (
	// Version contains the application version number. It's set via ldflags
	// when building.
	Version = ""

	// CommitSHA contains the SHA of the commit that this application was built
	// against. It's set via lgflags when building.
	CommitSHA = ""

	version = flag.Bool("version", false, "display version")
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Soft Serve, a self-hostable Git server for the command line.\n\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	if *version {
		if len(CommitSHA) > 7 {
			CommitSHA = CommitSHA[:7]
		}
		if Version == "" {
			Version = "(built from source)"
		}

		fmt.Printf("Soft Serve %s", Version)
		if len(CommitSHA) > 0 {
			fmt.Printf(" (%s)", CommitSHA)
		}

		fmt.Println()
		os.Exit(0)
	}

	cfg := config.DefaultConfig()
	s := server.NewServer(cfg)
	log.Printf("Starting SSH server on %s:%d\n", cfg.Host, cfg.Port)
	err := s.Start()
	if err != nil {
		log.Fatalln(err)
	}
}