package brew

import (
	"github.com/cavaliergopher/grab/v3"
	"github.com/kedare/kedare-setup/internal/tools"
	log "github.com/sirupsen/logrus"
)

var (
	BREW_SETUP_TEMP_FILE = "/tmp/brew-install.sh"
	BREW_SETUP_URI       = "https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh"
	BREW_SETUP_COMMAND   = []string{BREW_SETUP_TEMP_FILE}
)

func setup() error {
	// Not working as it required to sudo during the script run, will find a way later

	log.Infof("Downloading homebrew setup file from %s", BREW_SETUP_URI)
	_, err := grab.Get(BREW_SETUP_TEMP_FILE, BREW_SETUP_URI)
	if err != nil {
		log.Error("Failed to download setup")
		log.Error(err)
		return err
	}
	log.Infof("Downloaded successfully to %s", BREW_SETUP_TEMP_FILE)

	log.Infof("Executing setup file")
	_, err = tools.ExecAndStream("bash", BREW_SETUP_COMMAND...)
	if err != nil {
		log.Error("Error while executing homebrew setup")
		log.Error(err)
		return err
	}

	return nil
}
