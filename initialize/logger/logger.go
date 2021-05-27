package logger

import (
	"os"
	"time"

	"go-vue-next-server/initialize/config"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Init Zapの初期化を行う
func Init(filename string) *zap.Logger {
	c := config.GetConfig()
	filepath := c.GetString("log.path") + "/" + filename + ".%Y-%m-%d"

	rotate, err := rotatelogs.New(
		filepath,
		rotatelogs.WithRotationTime(time.Hour),
	)
	if err != nil {
		panic(err)
	}

	w := zapcore.AddSync(rotate)
	// &lumberjack.Logger{
	// 	Filename:   filename,
	// 	MaxSize:    c.GetInt("log.maxsize"),
	// 	MaxBackups: c.GetInt("log.maxbackups"),
	// 	MaxAge:     c.GetInt("log.maxage"),
	// },
	// )

	e := zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "file",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(e),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), w),
		zap.InfoLevel,
	)

	//ログレベルエラーからはstacktraceを含める
	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel), zap.Development())
	return logger
}
