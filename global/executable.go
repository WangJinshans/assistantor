package global

import (
	"github.com/rs/zerolog/log"
	"os"
)

func GetExecutablePath() (executablePath string, err error) {
	executablePath, err = os.Getwd()
	if err != nil {
		log.Info().Msgf("fail to get base path, error is: %v", err)
		return
	}
	return
}
