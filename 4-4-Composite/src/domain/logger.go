package domain

import (
	"C4M4H1/domain/enum"
	"errors"
)

type Logger struct {
	name     string
	level    *enum.Level
	exporter Exporter
	layout   Layout
	parent   *Logger
}

func newLogger(level *enum.Level, parent *Logger, name string, exporter Exporter, layout Layout) (*Logger, error) {
	if parent == nil {
		return nil, errors.New("parent cannot be nil")
	}
	return &Logger{name, level, exporter, layout, parent}, nil
}

func newRootLogger(level *enum.Level, exporter Exporter, layout Layout) (*Logger, error) {
	if exporter == nil {
		return nil, errors.New("exporter cannot be nil")
	}

	if layout == nil {
		return nil, errors.New("layout cannot be nil")
	}

	return &Logger{"Root", level, exporter, layout, nil}, nil
}

func (l *Logger) isRoot() bool {
	return l.parent == nil
}

func (l *Logger) getLevel() enum.Level {
	if l.isRoot() || l.level != nil {
		return *l.level
	}
	return l.parent.getLevel()
}

func (l *Logger) getExporter() Exporter {
	if l.isRoot() || l.exporter != nil {
		return l.exporter
	}

	return l.parent.getExporter()
}

func (l *Logger) getLayout() Layout {
	if l.isRoot() || l.layout != nil {
		return l.layout
	}
	return l.parent.getLayout()
}

func log(l *Logger, level enum.Level, message string) {
	if level >= l.getLevel() {
		output := l.getLayout().output(message, level, l)
		l.getExporter().export(output)
	}
}

func (l *Logger) Trace(message string) {
	log(l, enum.TRACE, message)
}

func (l *Logger) Info(message string) {
	log(l, enum.INFO, message)
}

func (l *Logger) Debug(message string) {
	log(l, enum.DEBUG, message)
}

func (l *Logger) Warn(message string) {
	log(l, enum.WARN, message)
}

func (l *Logger) Error(message string) {
	log(l, enum.ERROR, message)
}
