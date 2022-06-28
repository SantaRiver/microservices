package main

import "ozon/internal/payment/internal/server"

func main() {
	go server.RunRest()
	server.RunGrpc()
}
