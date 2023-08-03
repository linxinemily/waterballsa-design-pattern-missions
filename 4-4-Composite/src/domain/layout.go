package domain

import "C4M4H1/domain/enum"

type Layout interface {
	output(message string, level enum.Level, logger *Logger) string
}
