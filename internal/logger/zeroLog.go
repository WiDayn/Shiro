package logger

import "github.com/rs/zerolog"

func InitZeroLog() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
}
