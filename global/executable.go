package global

import (
	"github.com/rs/zerolog/log"
	"os"
	"path/filepath"
)

func GetExecutablePath() (executablePath string, err error) {
	executablePath, err = os.Getwd()
	if err != nil {
		log.Info().Msgf("fail to get base path, error is: %v", err)
		return
	}
	return
}

func GetReportPath() (filePath string, err error) {
	var executablePath string
	executablePath, err = os.Getwd()
	if err != nil {
		log.Info().Msgf("fail to get base path, error is: %v", err)
		return
	}
	filePath = filepath.Join(executablePath, "static", "report")
	return
}
