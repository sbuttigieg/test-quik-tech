package config

import (
	"io"
	"os"

	"github.com/sbuttigieg/test-quik-tech/wallet"
	"github.com/sirupsen/logrus"
)

func NewLogger(c *wallet.Config, f *os.File) *logrus.Logger {
	return &logrus.Logger{
		Out:       io.MultiWriter(f, os.Stdout),
		Level:     c.LogLevel,
		Formatter: &logrus.JSONFormatter{},
	}
}
