package asdf

import (
	"fmt"
	"github.com/kedare/kedare-setup/internal/tools"
	log "github.com/sirupsen/logrus"
	"os/user"
	"path/filepath"
)

var (
	ASDF_GITHUB_REPO = "https://github.com/asdf-vm/asdf.git"
	ASDF_SHELL       = "bash"
	ASDF_PATH        = ".asdf"
)

func getAsdfPath() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}

	return filepath.Join(usr.HomeDir, ASDF_PATH), nil
}

func setupPreCheck() error {
	return nil
}

func setup() error {
	log.Infof("Cloning ASDF repository")
	asdfPath, err := getAsdfPath()
	if err != nil {
		return err
	}
	_, err = tools.ExecAndStream(ASDF_SHELL, []string{"-c", fmt.Sprintf("git clone %s %s", ASDF_GITHUB_REPO, asdfPath)}...)
	if err != nil {
		log.Error("Error cloning ASDF repository")
		log.Error(err)
		return err
	}

	return nil
}
