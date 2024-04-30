package plugin

import (
	"fmt"
	"message-collect/pluginTemplate/scriptPlugin"
)

type EurBuildPlugin struct {
}

func (p EurBuildPlugin) StartTask() {
	resultCh := make(chan string)
	fmt.Println("启动fedora-messaging消费脚本")
	go scriptPlugin.ExecuteScript("/Users/shishupei/go/message-collect/script/fedora-messaging/consume_msg.bash")
	// Release the lock so other processes or threads can obtain a lock.
	<-resultCh
}
