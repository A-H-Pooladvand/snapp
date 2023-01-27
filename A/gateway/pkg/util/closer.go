package util

import (
	"gateway/pkg/log"
	"io"
)

// Closer handles closer interface in a one single API
// If we want to change it later we can simply change it all in here
func Closer(closer io.Closer) {
	if err := closer.Close(); err != nil {
		log.Error(err)
	}
}
