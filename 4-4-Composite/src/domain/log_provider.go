package domain

import (
	"C4M4H1/domain/enum"
	"encoding/json"
	"errors"
)

type LogProvider struct {
	loggers map[string]*Logger
}

func NewLogProvider() *LogProvider {
	loggers := make(map[string]*Logger)
	return &LogProvider{loggers}
}

func (p *LogProvider) GetLogger(name string) *Logger {
	logger, exists := p.loggers[name]
	if exists == false {
		return nil
	}
	return logger
}

func (p *LogProvider) CreateLogger(level *enum.Level, parent *Logger, name string, exporter Exporter, layout Layout) (*Logger, error) {
	if p.GetLogger(name) != nil {
		return nil, errors.New("logger already exists")
	}
	logger, err := newLogger(level, parent, name, exporter, layout)
	if err != nil {
		return nil, err
	}
	p.loggers[name] = logger
	return p.loggers[name], nil
}

func (p *LogProvider) CreateRootLogger(level *enum.Level, exporter Exporter, layout Layout) (*Logger, error) {
	if p.GetLogger("Root") != nil {
		panic("root logger already exists.")
	}
	logger, err := newRootLogger(level, exporter, layout)
	if err != nil {
		return nil, err
	}
	p.loggers["Root"] = logger
	return p.loggers["Root"], nil
}

func (p *LogProvider) ConfigureFromJSON(configData []byte) error {
	// 解析配置檔案
	config := LoggerConfig{}
	err := json.Unmarshal(configData, &config)
	if err != nil {
		return err
	}

	_, err = createLogger(p, config, nil)
	if err != nil {
		return err
	}

	return nil
}

type LoggerConfig struct {
	Name           string         `json:"name,omitempty"`
	LevelThreshold string         `json:"levelThreshold"`
	Exporter       ExporterConfig `json:"exporter"`
	Layout         string         `json:"layout"`
	Children       []LoggerConfig `json:"children"`
}

type ExporterConfig struct {
	Type     string           `json:"type"`
	FileName string           `json:"filename,omitempty"`
	Children []ExporterConfig `json:"children"`
}

func createLogger(p *LogProvider, config LoggerConfig, parent *Logger) (*Logger, error) {
	logLevel := enum.ParseLevel(config.LevelThreshold)
	var logger *Logger
	var err error
	if parent == nil {
		logger, err = p.CreateRootLogger(&logLevel, createExporter(config.Exporter), createLayout(config.Layout))
	} else {
		logger, err = p.CreateLogger(&logLevel, parent, config.Name, createExporter(config.Exporter), createLayout(config.Layout))
	}

	if err != nil {
		return nil, err
	}

	for _, child := range config.Children {
		_, err = createLogger(p, child, logger)
		if err != nil {
			return nil, err
		}
	}

	return logger, err
}

func createExporter(config ExporterConfig) Exporter {
	switch config.Type {
	case "console":
		return NewConsoleExporter()
	case "file":
		return NewFileExporter(config.FileName)
	case "composite":
		children := make([]Exporter, len(config.Children))
		for i, childConfig := range config.Children {
			children[i] = createExporter(childConfig)
		}
		return NewCompositeExporter(children)
	default:
		return nil
	}
}

func createLayout(layoutType string) Layout {
	switch layoutType {
	case "standard":
		return NewStandardLayout()
	default:
		return nil
	}
}
