package main

import (
	"flag"
	"github.com/kedare/kedare-setup/internal/lua"
	log "github.com/sirupsen/logrus"
)

const (
	VERSION = "0.1"
)

var (
	luaFile   = flag.String("config", "", "The configuration file to use")
	debugMode = flag.Bool("debug", false, "Enable debug mode")
)

func main() {
	flag.Parse()
	if *debugMode {
		log.SetLevel(log.DebugLevel)
		log.SetReportCaller(true)
	}
	log.Infof("WKST WorK Setup v%v", VERSION)
	log.Info("This will configure the system following a given pattern")

	if *luaFile == "" {
		log.Fatalf("No configuration file provided")
	}
	log.Infof("Loading configuration from %s", *luaFile)

	lua.RunLuaEngine(*luaFile)
}
