package logger

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	// Log is global logger
	Log *zap.Logger

	// ZLevel is global logger
	ZLevel *zap.AtomicLevel

	// timeFormat is custom Time format
	customTimeFormat string

	// onceInit guarantee initialize logger only once
	onceInit sync.Once

	// MSName from config
	MSName string
)

// getZapLevel level
func getZapLevel(level string) zapcore.Level {
	switch level {
	case "info", "INFO":
		return zapcore.InfoLevel
	case "warn", "WARN":
		return zapcore.WarnLevel
	case "debug", "DEBUG":
		return zapcore.DebugLevel
	case "error", "ERROR":
		return zapcore.ErrorLevel
	case "fatal", "FATAL":
		return zapcore.FatalLevel
	case "panic", "PANIC":
		return zapcore.PanicLevel
	case "dpanic", "DPANIC":
		return zapcore.DPanicLevel
	default:
		return zapcore.InfoLevel
	}
}

// customTimeEncoder encode Time to our custom format
// This example how we can customize zap default functionality
func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(time.RFC3339))
}

// Init initializes log by input parameters
// lvl - global log level: Debug(-1), Info(0), Warn(1), Error(2), DPanic(3), Panic(4), Fatal(5)
// timeFormat - custom time format for logger of empty string to use default
func Init(lvl string, timeFormat string, msName string) error {
	var err error

	onceInit.Do(func() {

		// Get MicroserviceName
		MSName = msName

		// First, define our level-handling logic.
		globalLevel := getZapLevel(lvl)

		// High-priority output should also go to standard error, and low-priority
		// output should also go to standard out.
		// It is usefull for Kubernetes deployment.
		// Kubernetes interprets os.Stdout log items as INFO and os.Stderr log items
		// as ERROR by default.
		highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl >= zapcore.ErrorLevel
		})
		lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl >= globalLevel && lvl < zapcore.ErrorLevel
		})
		consoleInfos := zapcore.Lock(os.Stdout)
		consoleErrors := zapcore.Lock(os.Stderr)

		// Configure console output.
		var useCustomTimeFormat bool
		ecfg := zap.NewProductionEncoderConfig()
		if len(timeFormat) > 0 {
			customTimeFormat = timeFormat
			ecfg.EncodeTime = customTimeEncoder
			useCustomTimeFormat = true
		}
		consoleEncoder := zapcore.NewJSONEncoder(ecfg)

		// Join the outputs, encoders, and level-handling functions into
		// zapcore.
		core := zapcore.NewTee(
			zapcore.NewCore(consoleEncoder, consoleErrors, highPriority),
			zapcore.NewCore(consoleEncoder, consoleInfos, lowPriority),
		)

		// From a zapcore.Core, it's easy to construct a Logger.
		Log = zap.New(core)
		zap.RedirectStdLog(Log)

		// If timeformat is not provided, log a warning on console to use a zap default or custom one
		if !useCustomTimeFormat {
			Log.Warn("time format for logger is not provided - use zap default")
		}
	})

	return err
}

// AppLogger func to be used for application logging
func AppLogger(level string, description string, time int64, items ...string) {

	var fields []zapcore.Field
	var field zapcore.Field
	if level != "info" {
		fields = append(fields, zap.Float64("elapsed-ms", float64(time)/1000000.0))
	}

	for i := range items {
		if strings.Contains(items[i], "|") {
			item := strings.Split(items[i], "|")
			field = zap.String(item[0], item[1])
			fields = append(fields, field)
		} else {
			fmt.Printf("Logger message item %v don't have separater '|' for including message title, description!\n", items[i])
		}
	}

	switch level {
	case "info":
		Log.Info(description, fields...)
	case "debug":
		Log.Debug(description, fields...)
	case "error":
		Log.Error(description, fields...)
	case "fatal":
		Log.Fatal(description, fields...)
	case "warn":
		Log.Warn(description, fields...)
	}
}
