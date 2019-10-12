package controller

import (
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var allowedMediaType = []string{"video/mp4"}

type ErrorInfo struct {
	Time  time.Time
	Path  string
	Method string
	Message string
}

func GetAllowFormat(format string, allowFormat []string) string {
	var allow string
	for _, item := range allowFormat {
		if format == item {
			allow = item
		}
	}
	return allow
}

func(control *Controller) LogError(c *gin.Context, message string) {
	var errorInfo ErrorInfo
	errorInfo.Time = time.Now()
	errorInfo.Path = c.FullPath()
	errorInfo.Method = c.Request.Method
	errorInfo.Message = message

	err := control.DB.Create(errorInfo).Error
	if err != nil {
		var log = logrus.New()
		log.Out = os.Stdout
		file, err := os.OpenFile("log/logging.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
		if err == nil {
			log.Out = file
		} else {
			log.Println(err)
		}

		log.WithFields(logrus.Fields{
			"Path":    errorInfo.Path,
			"Method": errorInfo.Method,
		}).Error(message)
	}
}
