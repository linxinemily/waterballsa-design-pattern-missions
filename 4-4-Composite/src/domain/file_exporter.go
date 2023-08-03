package domain

import (
	"os"
)

type FileExporter struct {
	filename string
}

func NewFileExporter(filename string) *FileExporter {
	return &FileExporter{filename}
}

func (e *FileExporter) export(message string) error {
	logFile := e.filename

	// 檢查檔案是否存在，不存在則創建
	if _, err := os.Stat(logFile); os.IsNotExist(err) {
		_, err := os.Create(logFile)
		if err != nil {
			return err
		}
	}

	// 開啟檔案以寫入模式
	file, err := os.OpenFile(logFile, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// 檢查最後一個字元是否為換行符號，如果不是則添加換行符號
	if message[len(message)-1] != '\n' {
		message += "\n"
	}

	// 寫入日誌
	_, err = file.WriteString(message)
	if err != nil {
		return err
	}
	return nil
}
