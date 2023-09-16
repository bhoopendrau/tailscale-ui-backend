package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/bhoopendrau/tailscale-ui-backend/config"
	"github.com/bhoopendrau/tailscale-ui-backend/db"
	"github.com/bhoopendrau/tailscale-ui-backend/server"
)

func main() {
	environment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*environment)
	db.Init()
	server.Init()
}
