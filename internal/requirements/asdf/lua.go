package asdf

import (
	log "github.com/sirupsen/logrus"
	lua "github.com/yuin/gopher-lua"
)

var exports = map[string]lua.LGFunction{
	"setup":           luaSetup,
	"ensure_packages": luaEnsurePackages,
	"ensure_plugins":  luaEnsurePlugins,
}

func luaSetup(L *lua.LState) int {
	setup()
	return 0
}

func luaEnsurePackages(L *lua.LState) int {
	log.Error("Not implemented yet")
	return 1
}

func luaEnsurePlugins(L *lua.LState) int {
	log.Error("Not implemented yet")
	return 1
}

func LuaLoader(L *lua.LState) int {
	mod := L.SetFuncs(L.NewTable(), exports)
	L.Push(mod)
	return 1
}
