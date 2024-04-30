package manager

import (
	"message-collect/pluginTemplate/scriptPlugin"
)

var (
	scriptPluginManager *ScriptPluginManager
)

type ScriptPluginManager struct {
	plugins map[string]scriptPlugin.ScriptPlugin
}

func init() {
	scriptPluginManager = NewScriptPluginManager()
}

func NewScriptPluginManager() *ScriptPluginManager {
	return &ScriptPluginManager{
		plugins: make(map[string]scriptPlugin.ScriptPlugin),
	}
}

func (pm *ScriptPluginManager) Register(name string, plugin scriptPlugin.ScriptPlugin) {
	pm.plugins[name] = plugin
}

func (pm *ScriptPluginManager) Get(name string) scriptPlugin.ScriptPlugin {
	plugin, ok := pm.plugins[name]
	if !ok {
		return nil
	}
	return plugin
}

func StartTask(scriptPlugin scriptPlugin.ScriptPlugin) {
	scriptPluginManager.Register("A", scriptPlugin)
	pluginA := scriptPluginManager.Get("A")
	if pluginA != nil {
		go pluginA.StartTask()
	}
}
