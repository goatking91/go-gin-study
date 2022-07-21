package logger

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	S *zap.SugaredLogger // zap sugared logger
	L *zap.Logger        // zap logger
)

func init() {
	logLevel := parseLevel(os.Getenv("LOG_LEVEL"))

	cf := zap.NewProductionEncoderConfig()
	cf.EncodeTime = zapcore.ISO8601TimeEncoder // ISO8601-formatted (2022-05-06T17:14:21.101+0900) string with millisecond precision
	cEncoder := zapcore.NewConsoleEncoder(cf)
	var core zapcore.Core

	// 로그파일에 출력할 경우 파일과 콘솔출력
	if s, err := strconv.ParseBool(os.Getenv("LOG_WRITE_FILE")); err == nil && s {
		fEncoder := zapcore.NewConsoleEncoder(cf) // NewJSONEncoder(cf)
		logFile := os.Getenv("LOG_PATH") + "/" + os.Getenv("LOG_FILE_NAME")

		duration, err := time.ParseDuration(os.Getenv("LOG_ROTATE_MAX_AGE"))
		if err != nil {
			log.Printf("error setup logging. %v", err)
		}

		// set logrotate
		logf, err := rotatelogs.New(
			logFile+"."+os.Getenv("LOG_ROTATE_PATTERN"),
			rotatelogs.WithLinkName(logFile),       // 날짜가 없는 파일명으로 링크 생성
			rotatelogs.WithMaxAge(duration),        // 보관일
			rotatelogs.WithRotationTime(time.Hour)) // 시간당 로테이트 동작
		if err != nil {
			log.Printf("error setup logging. %v", err)
		}

		fWriter := zapcore.AddSync(logf)

		core = zapcore.NewTee(
			zapcore.NewCore(fEncoder, fWriter, logLevel),
			zapcore.NewCore(cEncoder, zapcore.AddSync(os.Stdout), logLevel),
		)
	} else {
		// 콘솔만 출력하는 경우 (Docker 로 실행시 유용함)
		core = zapcore.NewTee(
			zapcore.NewCore(cEncoder, zapcore.AddSync(os.Stdout), logLevel),
		)
	}

	L = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	S = L.Sugar()
}

func parseLevel(lvl string) zapcore.Level {
	switch strings.ToUpper(lvl) {
	case "FATAL":
		return zapcore.FatalLevel
	case "ERROR":
		return zapcore.ErrorLevel
	case "WARN", "WARNING":
		return zapcore.WarnLevel
	case "INFO":
		return zapcore.InfoLevel
	case "DEBUG":
		return zapcore.DebugLevel
		//case "TRACE":
		//	return TRACE
	}

	log.Printf("not a valid log Level: %q", lvl)
	return zapcore.InfoLevel
}
