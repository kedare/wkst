package brew

import (
	"github.com/kedare/kedare-setup/internal/tools"
	log "github.com/sirupsen/logrus"
)

func ensurePackages(packages []string) error {
	log.Infof("Got a list of %v packages to ensure", len(packages))
	parameters := append([]string{"install"}, packages...)
	_, err := tools.ExecAndStream("brew", parameters...)
	if err != nil {
		return err
	}
	return nil
}
