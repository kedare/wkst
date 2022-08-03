package lua

import (
	"github.com/kedare/kedare-setup/internal/requirements/asdf"
	"github.com/kedare/kedare-setup/internal/requirements/brew"
	log "github.com/sirupsen/logrus"
	lua "github.com/yuin/gopher-lua"
)

var LUA_MODULES = map[string]func(state *lua.LState) int{
	"brew": brew.LuaLoader,
	"asdf": asdf.LuaLoader,
}

func RunLuaEngine(luaFile string) {
	log.Debug("Building Lua engine")
	L := lua.NewState()
	defer L.Close()

	// Load modules
	for module, impl := range LUA_MODULES {
		log.Debugf("Loading module %s", module)
		L.PreloadModule(module, impl)
	}

	log.Infof("Executing %s", luaFile)
	if err := L.DoFile(luaFile); err != nil {
		log.Panic(err)
	}
	log.Debug("Lua execution completed without error")
}
