package manager

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/opensourceways/message-collect/pluginTemplate/webServerPlugin"
)

var (
	webServerPluginManager *WebServerPluginManager
	engine                 = gin.Default()
)

type WebServerPluginManager struct {
	plugins map[string]webServerPlugin.WebServerPlugin
}

func init() {
	webServerPluginManager = NewWebServerPluginManager()
}

func NewWebServerPluginManager() *WebServerPluginManager {
	return &WebServerPluginManager{
		plugins: make(map[string]webServerPlugin.WebServerPlugin),
	}
}

func (pm *WebServerPluginManager) Register(name string, plugin webServerPlugin.WebServerPlugin) {
	pm.plugins[name] = plugin
}

func (pm *WebServerPluginManager) Get(name string) webServerPlugin.WebServerPlugin {
	plugin, ok := pm.plugins[name]
	if !ok {
		return nil
	}
	return plugin
}

func AddRoute(webServerPlugin webServerPlugin.WebServerPlugin) {
	webServerPluginManager.Register("A", webServerPlugin)
	pluginA := webServerPluginManager.Get("A")
	if pluginA != nil {
		pluginA.AddRoute()
	}
}

func Run(port int) {
	engine.Run(fmt.Sprintf(":%d", port))
}
