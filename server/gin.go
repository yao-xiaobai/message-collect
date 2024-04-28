/*
Copyright (c) Huawei Technologies Co., Ltd. 2023. All rights reserved
*/

// Package server provides functionality for setting up and configuring a server for handling code repo operations.
package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"message-collect/server/gitee"
	"strings"
	"time"
)

// StartWebServer starts a web server with the given configuration.
// It initializes the services, sets up the routers for different APIs, and starts the server.
// If TLS key and certificate are provided, it will use HTTPS.
// If removeCfg is true, it will remove the key and certificate files after starting the server.
func StartWebServer(port int) {
	engine := gin.Default()
	addRoute(engine)
	engine.Run(fmt.Sprintf(":%d", port))
}

func addRoute(engine *gin.Engine) {
	gitee.CommentWebhook(engine)
}

func logRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		c.Next()

		endTime := time.Now()

		errmsg := ""
		for _, ginErr := range c.Errors {
			if errmsg != "" {
				errmsg += ","
			}
			errmsg = fmt.Sprintf("%s%s", errmsg, ginErr.Error())
		}

		if strings.Contains(c.Request.RequestURI, "/swagger/") ||
			strings.Contains(c.Request.RequestURI, "/internal/heartbeat") {
			return
		}

		log := fmt.Sprintf(
			"| %d | %d | %s | %s ",
			c.Writer.Status(),
			endTime.Sub(startTime),
			c.Request.Method,
			c.Request.RequestURI,
		)
		if errmsg != "" {
			log += fmt.Sprintf("| %s ", errmsg)
		}

		logrus.Info(log)
	}
}
