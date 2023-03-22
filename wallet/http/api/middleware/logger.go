package middleware

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/sbuttigieg/test-quik-tech/wallet/models"
	"github.com/sirupsen/logrus"
)

func Logger(ctx context.Context, log *logrus.Logger) gin.HandlerFunc {
	return gin.LoggerWithFormatter(
		func(params gin.LogFormatterParams) string {
			logData := models.LogReqAPI{
				Method:     params.Method,
				Path:       params.Path,
				StatusCode: params.StatusCode,
			}
			log.Info(logData)

			return ""
		},
	)
}
