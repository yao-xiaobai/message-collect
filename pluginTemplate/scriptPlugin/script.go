package scriptPlugin

import (
	"fmt"
	"os/exec"
)

func ExecuteScript(scriptPath string) {
	// 使用exec.Command创建一个新的cmd对象，指定要执行的Shell命令
	cmd := exec.Command("sh", scriptPath)

	// 执行命令并捕获输出和错误
	err := cmd.Start()
	if err != nil {
		fmt.Println("Error starting script:", err)
		return
	}

	fmt.Println("Python script started asynchronously.")
	err = cmd.Wait()
	if err != nil {
		fmt.Println("Python script execution failed:", err)
	}
}
