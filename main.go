package main

import (
	"fmt"
	"github.com/bamzi/jobrunner"
	"github.com/go-redsync/redsync/v4"
	redislock "message-collect/common/redis"
	"message-collect/job"
	"time"
)

func main() {
	redislock.Init()
	jobrunner.Start() // optional: jobrunner.Start(pool int, concurrent int) (10, 1)
	fmt.Println("Main goroutine continues")
	startFedoraMsg()
	select {}

}

func startFedoraMsg() {
	mutex := redislock.RS.NewMutex("fed-msg", redsync.WithExpiry(time.Second*60))
	err := mutex.Lock()
	if err != nil {
		fmt.Println("Failed to acquire lock:", err)
		return
	}
	jobrunner.Now(
		job.FedoraMsg{},
	)
	// 主goroutine继续执行其他操作
}
