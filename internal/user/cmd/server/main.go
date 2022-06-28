package main

import "ozon/internal/user/internal/server"

func main() {
	go server.RunRest()
	server.RunGrpc()
}
