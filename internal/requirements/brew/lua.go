package brew

import (
	log "github.com/sirupsen/logrus"
	lua "github.com/yuin/gopher-lua"
)

var exports = map[string]lua.LGFunction{
	"setup":           luaSetup,
	"ensure_packages": luaEnsurePackages,
}

func luaSetup(L *lua.LState) int {
	log.Error("Not implemented yet")
	return 0
}

func luaEnsurePackages(L *lua.LState) int {
	log.Debug("Building package list from lua table")
	packages := []string{}
	luaPackages := L.CheckTable(2)
	luaPackages.ForEach(func(_, value lua.LValue) {
		switch value.(type) {
		case lua.LString:
			packages = append(packages, value.String())
		default:
			log.Errorf("%s is a %s, packages names need to be strings", value.String(), value.Type().String())
		}
	})
	log.Debugf("Got package list %s", packages)
	ensurePackages(packages)
	return 1
}

func LuaLoader(L *lua.LState) int {
	mod := L.SetFuncs(L.NewTable(), exports)
	L.Push(mod)
	return 1
}
