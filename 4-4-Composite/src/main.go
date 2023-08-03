package main

import (
	"C4M4H1/domain"
	"os"
)

func main() {

	provider := domain.NewLogProvider()
	configData, err := os.ReadFile("config.json")
	if err != nil {
		panic(err)
	}
	err = provider.ConfigureFromJSON(configData)
	if err != nil {
		panic(err)
	}

	rootLogger := provider.GetLogger("Root")
	rootLogger.Info("This is an info message")
	rootLogger.Debug("This is a debug message.")
	appGameLogger := provider.GetLogger("app.game")
	if appGameLogger != nil {
		appGameLogger.Info("This is an info message from app.game logger.")
		appGameAILogger := provider.GetLogger("app.game.ai")
		if appGameAILogger != nil {
			appGameAILogger.Trace("This is a trace message from app.game.ai logger.")
		}
	}
	//rootLoggerLevel := enum.DEBUG
	//rootLogger, err := provider.CreateRootLogger(
	//	&rootLoggerLevel,
	//	domain.NewCompositeExporter([]domain.Exporter{
	//		domain.NewConsoleExporter(),
	//		domain.NewFileExporter("daily.log"),
	//	}),
	//	domain.NewStandardLayout(),
	//)
	//if err != nil {
	//	panic(err)
	//}
	//
	//rootLogger.Info("父 logger")
	//
	//appLoggerLevel := enum.ERROR
	//appLogger, err := provider.CreateLogger(&appLoggerLevel, rootLogger, "app", nil, nil)
	//if err != nil {
	//	panic(err)
	//}
	//appLogger.Info("子 logger")
}
