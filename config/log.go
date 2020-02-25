package config

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/koushamad/goro-core/app/helper"
	"time"
)

var (
	// file console
	Log = map[string]string{
		"driver": helper.Env("LOG_DRIVER", "file"),
	}
)

func LogFormat(param gin.LogFormatterParams) string {
	return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
		param.ClientIP,
		param.TimeStamp.Format(time.RFC3339),
		param.Method,
		param.Path,
		param.Request.Proto,
		param.StatusCode,
		param.Latency,
		param.Request.UserAgent(),
		param.ErrorMessage,
	)
}
