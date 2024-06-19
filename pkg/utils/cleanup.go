package utils

import (
	"os"

	"github.com/Ns2Kracy/IceWhale-ZimaCube-Metrics/config"
)

func CleanUp() {
	// clean db
	path := config.AppInfo.DBPath + "metrics.db"
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return
		}

		os.Remove(path)
	}
}
