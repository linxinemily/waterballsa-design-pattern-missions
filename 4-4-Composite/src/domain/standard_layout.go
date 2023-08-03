package domain

import (
	"C4M4H1/domain/enum"
	"fmt"
	"time"
)

type StandardLayout struct {
}

func NewStandardLayout() *StandardLayout {
	return &StandardLayout{}
}

func (s *StandardLayout) output(message string, level enum.Level, logger *Logger) string {
	return fmt.Sprintf(
		"%s |-%s %s - %s",
		time.Now().Format("2006-01-02 15:04:05.000"),
		level,
		logger.name,
		message,
	)
}
