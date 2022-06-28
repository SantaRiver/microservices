package main

import "ozon/internal/order/internal/server"

func main() {
	go server.RunRest()
	server.RunGrpc()
}
