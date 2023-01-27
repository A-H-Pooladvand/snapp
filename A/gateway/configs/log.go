package configs

import (
	"fmt"
	"github.com/caarlos0/env/v6"
	"strings"
	"time"
)

var storagePath = "storage/logs"

type Log struct {
	Level string `env:"LOG_LEVEL" envDefault:"debug"`
}

func (l *Log) GetPath(path string) string {
	now := time.Now().Format("2006-01-02")

	paths := map[string]string{
		"default": fmt.Sprintf("%s/%s.log", storagePath, now),
	}

	// if path is nil we just return the default log filename
	if path == "" {
		path = paths["default"]
	}

	// if path is a key inside paths we will return the value
	if path, ok := paths[path]; ok {
		return path
	}

	// if the path contains storage/logs we don't change it
	ok := strings.Contains(path, storagePath)

	if ok {
		return path
	}

	// at last, we will add storage/logs to the path
	path = strings.TrimPrefix(path, "/")

	return fmt.Sprintf("%s/%s", storagePath, path)
}

func NewLog() Log {
	c := Log{}

	env.Parse(&c)

	return c
}
