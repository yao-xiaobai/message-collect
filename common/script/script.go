package script

import (
	"fmt"
	"os/exec"
	"sync"
)

func ExecuteScript(scriptPath string, wg *sync.WaitGroup) {
	defer wg.Done()

	// 使用exec.Command创建一个新的cmd对象，指定要执行的Shell命令
	cmd := exec.Command("sh", scriptPath)

	// 执行命令并捕获输出和错误
	output, err := cmd.CombinedOutput()
	if err != nil {
		// 如果执行过程中出现错误，输出错误信息
		fmt.Printf("Error executing script %s: %v\n", scriptPath, err)
	} else {
		// 输出Shell命令的执行结果
		fmt.Printf("Output of script %s:\n%s\n", scriptPath, string(output))
	}
}
