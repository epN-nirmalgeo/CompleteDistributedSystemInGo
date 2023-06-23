package main

import (
	"RedisGoLang/config"
	"RedisGoLang/server"
	"flag"
	"log"
)

func setupServer() {
	var host string
	var port int

	flag.StringVar(&host, "host", "0.0.0.0", "RedisGoLang Host")
	flag.IntVar(&port, "port", 6379, "RedisGoLang Port")
	flag.Parse()

	serverConfig := config.ServerConfig{
		ServerHost: host,
		ServerPort: port,
	}
	server.RunSyncTCPServer(serverConfig)
}

func main() {
	log.Println("Setting up RedisGoLang Server")
	setupServer()
}
