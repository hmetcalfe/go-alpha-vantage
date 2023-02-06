package utils

import (
	"encoding/base64"
	"fmt"

	log "github.com/sirupsen/logrus"
)

func DecodeAPIKey(key string) (string, error) {
	logger := log.WithField("function", "utils.DecodeAPIKey()")
	var retKey string

	// Allocate the appropriate size buffer
	buffer := make([]byte, base64.StdEncoding.DecodedLen(len(key)))
	n, err := base64.StdEncoding.Decode(buffer, []byte(key))

	if err != nil {
		logText := fmt.Sprintf("failed to decode key: %s", key)
		logger.WithError(err).Error(logText)

		return "", fmt.Errorf("%s: %w", logText, err)
	}

	retKey = string(buffer[:n])

	return retKey, nil
}
