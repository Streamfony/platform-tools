package logger

import (
	"time"

	"github.com/Streamfony/platform-tools/env"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	gormLog "gorm.io/gorm/logger"
	"moul.io/zapgorm2"
)

func NewLogger(options ...zap.Option) (*zap.Logger, error) {
	if env.IsDev() {
		return zap.NewDevelopment(options...)
	}

	return zap.NewProduction(options...)
}

func NewGinLoggerMiddleware(logger *zap.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ginzap.Ginzap(logger, time.RFC3339, true)(ctx)
		ginzap.RecoveryWithZap(logger, true)(ctx)
	}
}

func NewGormLogger(logger *zap.Logger) gormLog.Interface {
	if env.IsDev() {
		return zapgorm2.New(logger).LogMode(gormLog.Info)
	}

	gormLogger := zapgorm2.New(logger)
	gormLogger.IgnoreRecordNotFoundError = true
	return gormLogger.LogMode(gormLog.Error)
}
