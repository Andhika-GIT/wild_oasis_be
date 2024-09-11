package config

import (
	"fmt"
	"os"
	"time"

	"gorm.io/gorm/logger"
)

// untuk stdout logger pada &gorm.config{}

type logrusWriter struct {
	Writer *os.File
}

func (l *logrusWriter) Printf(message string, args ...interface{}) {
	fmt.Fprintf(l.Writer, message, args...)
}

func NewLogger() logger.Interface {
	
	writer := &logrusWriter{Writer: os.Stdout}

	newLogger := logger.New(
		writer,
		logger.Config{
			SlowThreshold:             time.Second * 2, // Threshold waktu untuk query lambat
			LogLevel:                  logger.Info,     // Level log (Info/Error/Warn/Silent)
			IgnoreRecordNotFoundError: true,            // Tidak log error untuk record not found
			Colorful:                  false,           // Nonaktifkan warna (berguna untuk log di file)
		},
	)

	return newLogger
}