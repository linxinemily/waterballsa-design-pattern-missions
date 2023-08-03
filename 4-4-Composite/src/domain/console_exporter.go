package domain

import "fmt"

type ConsoleExporter struct {
}

func NewConsoleExporter() *ConsoleExporter {
	return &ConsoleExporter{}
}

func (c *ConsoleExporter) export(message string) error {
	fmt.Println(message)
	return nil
}
