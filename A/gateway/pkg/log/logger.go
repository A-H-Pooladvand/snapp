package log

import (
	"gateway/configs"
	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
	"log"
)

type Log struct {
	logger *logrus.Logger
	config configs.Log
}

func Logger(v ...string) *logrus.Logger {
	l := new(Log)
	l.config = configs.NewLog()
	l.logger = logrus.New()
	l.logger.SetLevel(l.Level())
	l.logger.SetFormatter(&logrus.JSONFormatter{})

	dest := ""
	if len(v) > 0 {
		dest = v[0]
	}

	filename := l.config.GetPath(dest)

	output := l.OutputFile(filename)
	l.logger.SetOutput(output)

	return l.logger
}

func (l *Log) Level() logrus.Level {
	level, err := logrus.ParseLevel(
		l.config.Level,
	)

	if err != nil {
		log.Println(err)
	}

	return level
}

// OutputFile Indicates Logger output file
func (l *Log) OutputFile(filename string) *lumberjack.Logger {
	return &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    1,     // MegaBytes
		MaxBackups: 0,     // Pooladvand: For some reason if you set it to a value more than 0 the whole log system will fail, and it doesn't create more than 1 file
		MaxAge:     30,    // Days
		Compress:   false, // Pooladvand: Since we're using filebeat to ship the logs to logstash, so we will not compress log files.
	}
}

func Error(args ...any) {
	Logger().Error(args)
}
