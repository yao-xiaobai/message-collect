package job

import (
	"fmt"
	"message-collect/common/script"
)

type FedoraMsg struct {
}

// FedoraMes.Run() will get triggered automatically.
func (e FedoraMsg) Run() {
	resultCh := make(chan string)
	fmt.Println("启动fedora-messaging消费脚本")
	script.ExecuteScript("script/fedora-messaging/consume_msg.bash")
	// Release the lock so other processes or threads can obtain a lock.
	<-resultCh
}
