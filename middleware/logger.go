package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	rologs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"time"
)

func Logger() gin.HandlerFunc {
	filePath := "log/log"
	linkName:="latest_log.log"
	logger := logrus.New()

	logger.SetLevel(logrus.DebugLevel)

	logWriter, err := rologs.New(
		filePath+"%Y%m%d.log",
		rologs.WithMaxAge(7*24*time.Hour),
		rologs.WithRotationTime(24*time.Hour),
		rologs.WithLinkName(linkName),

	)
	if err != nil {
		log.Println(err)
	}
	writeMap := lfshook.WriterMap{
		logrus.DebugLevel: logWriter,
		logrus.InfoLevel:  logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
	}

	logger.AddHook(lfshook.NewHook(writeMap, &logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	}))

	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		endTime := time.Since(startTime)
		spendTime := fmt.Sprintf("%d us", endTime.Microseconds())
		hostName, err := os.Hostname()
		if err != nil {
			hostName = "anonymous host name"
		}
		statusCode := c.Writer.Status()
		clientIp := c.ClientIP()
		userAgent := c.Request.UserAgent()
		method := c.Request.Method
		path := c.Request.RequestURI

		entry := logger.WithFields(logrus.Fields{
			"status":    statusCode,
			"host":      hostName,
			"spendTime": spendTime,
			"ip":        clientIp,
			"method":    method,
			"userAgent": userAgent,
			"path":      path,
		})

		if len(c.Errors) > 0 {
			entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
		} else if statusCode > 500 {
			entry.Error()
		} else if statusCode > 400 {
			entry.Warn()
		} else {
			entry.Info()
		}
	}
}
