package server

import "github.com/bhoopendrau/tailscale-ui-backend/config"

func Init() {
	config := config.GetConfig()
	r := NewRouter()
	r.Run(config.GetString("server.port"))
}
