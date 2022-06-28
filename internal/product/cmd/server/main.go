package main

import "ozon/internal/product/internal/server"

func main() {
	go server.RunRest()
	server.RunGrpc()
}
