package oslog

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/OpenSlides/openslides-go/environment"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var isDev bool

// InitLog has to be called at startup to set the log format.
func InitLog(lookup environment.Environmenter) {
	devmode, _ := strconv.ParseBool(environment.EnvDevelopment.Value(lookup))

	cw := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: "2006-01-02T15:04:05:000",
		FormatLevel: func(i any) string {
			level := strings.ToUpper(fmt.Sprintf("%-6s", i))
			if !isDev {
				return level
			}
			switch i {
			case "debug":
				return fmt.Sprintf("\x1b[0m%s\x1b[0m", level)
			case "info":
				return fmt.Sprintf("\x1b[32m%s\x1b[0m", level) // Grün
			case "warn":
				return fmt.Sprintf("\x1b[33m%s\x1b[0m", level) // Gelb
			case "error":
				return fmt.Sprintf("\x1b[31m%s\x1b[0m", level) // Rot
			case "fatal", "panic":
				return fmt.Sprintf("\x1b[35m%s\x1b[0m", level) // Magenta
			default:
				return level
			}
		},
	}

	if devmode {
		isDev = true
		cw.TimeFormat = "15:04:05"
		cw.Out = os.Stderr
	}

	log.Logger = log.Output(cw)
}

// Debug writes a message at the debug level.
func Debug(format string, a ...any) {
	msg(log.Debug(), format, a...)
}

// Info writes a message at the info level.
func Info(format string, a ...any) {
	msg(log.Info(), format, a...)
}

// Warn writes a message at the warn level.
func Warn(format string, a ...any) {
	msg(log.Warn(), format, a...)
}

// Error writes a message at the error level.
func Error(format string, a ...any) {
	msg(log.Error(), format, a...)
}

// Metric writes a metric at info level.
func Metric(metric json.RawMessage) {
	if isDev {
		log.Info().Msg(string(metric))
		return
	}
	log.Info().RawJSON("metric", metric).Msg("")
}

func msg(e *zerolog.Event, format string, a ...any) {
	if isDev {
		e.Msgf(format, a...)
		return
	}
	e.Str("msg", fmt.Sprintf(format, a...)).Msg("")
}
