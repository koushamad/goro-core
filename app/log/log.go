package log

import (
	"github.com/gin-gonic/gin"
	"github.com/koushamad/goro-app/app/helper"
	"github.com/koushamad/goro-core/app/exception/throw"
	"github.com/koushamad/goro-core/config"
	"io"
	"os"
)

func Logger() gin.HandlerFunc {
	FileLogger()
	return gin.LoggerWithFormatter(config.LogFormat)
}

func FileLogger() {
	gin.DisableConsoleColor()
	file, err := os.Open(helper.LogPath())
	if err != nil {
		file, err = os.Create(helper.LogPath())
	}
	throw.Fatal("Cannot open or create log file", 113, err)
	gin.DefaultWriter = io.MultiWriter(file, os.Stdout)
}
