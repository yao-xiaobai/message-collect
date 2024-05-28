package plugin

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	kfklib "github.com/opensourceways/kafka-lib/agent"
	serverModels "github.com/opensourceways/message-collect/plugin/models"
	"net/http"
)

type GiteeServerPlugin struct {
	Engine *gin.Engine
}

func (p GiteeServerPlugin) AddRoute() {
	p.Engine.POST("/webhook/gitee/comment", handleComment)
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
	msg, _ := json.Marshal(payload)
	kfklib.Publish("gitee_comment_raw", nil, msg)

	// 返回成功响应
	c.JSON(http.StatusOK, gin.H{"message": "Webhook received successfully"})
}
