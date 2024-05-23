package manager

import (
	"github.com/opensourceways/message-collect/pluginTemplate/mqPlugin"
	"reflect"
)

var (
	MqPluginManager *mqPluginManager
)

type mqPluginManager struct {
	plugins map[string]mqPlugin.MqPlugin
}

func init() {
	MqPluginManager = NewMqPluginManager()
}

func NewMqPluginManager() *mqPluginManager {
	return &mqPluginManager{
		plugins: make(map[string]mqPlugin.MqPlugin),
	}
}

func (pm *mqPluginManager) Register(name string, plugin mqPlugin.MqPlugin) {
	pm.plugins[name] = plugin
}

func (pm *mqPluginManager) Get(name string) mqPlugin.MqPlugin {
	plugin, ok := pm.plugins[name]
	if !ok {
		return nil
	}
	return plugin
}

func StartConsume(plugin mqPlugin.MqPlugin) {
	registerKey := reflect.TypeOf(plugin).Name()
	MqPluginManager.Register(registerKey, plugin)
	pluginA := MqPluginManager.Get(registerKey)
	if pluginA != nil {
		go pluginA.StartConsume()
	}
}
