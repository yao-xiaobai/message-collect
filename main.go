package main

import (
	"message-collect/common/script"
	"sync"
)

func main() {
	// 创建一个通道来通知任务完成
	var wg sync.WaitGroup
	wg.Add(1)
	// 异步执行脚本
	go script.ExecuteScript("script/fedora-messaging/consume_msg.bash", &wg)
	wg.Wait()
}
