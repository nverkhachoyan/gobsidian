package main

import (
	"flag"
	"os"

	"github.com/charmbracelet/log"
)

func main() {
	watchFlag := flag.Bool("watch", false, "Watch for file changes and rebuild automatically")
	clearFlag := flag.Bool("clear", false, "Clear the public folder before generating")
	serveFlag := flag.Bool("serve", false, "Serve the generated website after building")
	portFlag := flag.String("port", "8080", "Port to serve the website on")
	flag.Parse()

	app, err := NewApp()
	if err != nil {
		log.Error("Failed to initialize application", "error", err)
		os.Exit(1)
	}

	app.Run(*clearFlag, *watchFlag, *serveFlag, *portFlag)
}
