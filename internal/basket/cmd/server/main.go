package main

import "ozon/internal/basket/internal/server"

func main() {
	go server.RunRest()
	server.RunGrpc()
}
