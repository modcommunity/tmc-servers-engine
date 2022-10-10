package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/gamemann/tmc-servers-engine/internal/Config"
)

const HELP_MENU = "Help Options\n\t-cfg= --cfg -cfg <path> > Path to config file override.\n\t-l --list > Print out full config.\n\t-v --version > Print out version and exit.\n\t-h --help > Display help menu.\n\n"
const VERSION = "1.0.0"

func main() {
	var list bool
	var version bool
	var help bool

	// Setup simple flags (booleans).
	flag.BoolVar(&list, "list", false, "Print out config and exit.")
	flag.BoolVar(&list, "l", false, "Print out config and exit.")

	flag.BoolVar(&version, "version", false, "Print out version and exit.")
	flag.BoolVar(&version, "v", false, "Print out version and exit.")

	flag.BoolVar(&help, "help", false, "Print out help menu and exit.")
	flag.BoolVar(&help, "h", false, "Print out help menu and exit.")

	// Look for 'cfg' flag in command line arguments (default path: ./settings.json).
	configFile := flag.String("cfg", "settings.json", "The path to the Rust Auto Wipe config file.")

	// Parse flags.
	flag.Parse()

	// Check for version flag.
	if version {
		fmt.Print(VERSION)

		os.Exit(0)
	}

	// Check for help flag.
	if help {
		fmt.Print(HELP_MENU)

		os.Exit(0)
	}

	// Create config struct.
	cfg := Config.Config{}

	// Set config defaults.
	cfg.SetDefaults()

	// Attempt to read config.
	err := cfg.LoadConfig(*configFile)

	// If we have no config, create the file with the defaults.
	if err != nil {
		// If there's an error and it contains "no such file", try to create the file with defaults.
		if strings.Contains(err.Error(), "no such file") {
			err = cfg.WriteDefaultsToFile(*configFile)

			if err != nil {
				fmt.Println("Failed to open config file and cannot create file.")
				fmt.Println(err)

				os.Exit(1)
			}
		}

		fmt.Println("WARNING - No config file found. Created config file at " + *configFile + " with defaults.")
	}

	// Check for list flag.
	if list {
		// Encode config as JSON string.
		json_data, err := json.MarshalIndent(cfg, "", "   ")

		if err != nil {
			fmt.Println(err)

			os.Exit(1)
		}

		fmt.Println(string(json_data))

		os.Exit(0)
	}

	// Signal.
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM)
	<-sigc

	os.Exit(0)
}
