package gitee

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"message-collect/common/kafka"
	serverModels "message-collect/models/server"
	"net/http"
)

func CommentWebhook(engine *gin.Engine) {
	// Webhook 接口
	engine.POST("/webhook/gitee/comment", handleComment)
}

func handleComment(c *gin.Context) {
	// 读取请求体
	var payload serverModels.Hook
	if err := c.BindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse JSON data"})
		return
	}

	fmt.Println("Received Gitee comment event:")
	fmt.Println("Action:", payload.HookName)
	fmt.Println("Comment ID:", payload.Comment.ID)
	fmt.Println("Comment Body:", payload.Comment.Body)
	kafka.KfkProducer.SendMessage("gitee_comment_raw", payload)

	// 返回成功响应
	c.JSON(http.StatusOK, gin.H{"message": "Webhook received successfully"})
}
