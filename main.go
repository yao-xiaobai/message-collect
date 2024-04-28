package main

import (
	"fmt"
	"message-collect/common/kafka"
	redislock "message-collect/common/redis"
	"message-collect/script"
	"message-collect/server"
)

func main() {
	redislock.Init()
	kafka.Init()

	fmt.Println("Main goroutine continues")
	script.StartScript()
	server.StartWebServer(8081)
	select {}
}
