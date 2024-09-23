package io

import (
	"chat/core"
	"fmt"
	"os"
)

type TextIoService struct {
	outputFolder string
}

func NewTextService() core.IOService {
	return &TextIoService{
		outputFolder: "./build",
	}
}

func (io *TextIoService) Execute(fileName, fileContent string) error {
	fileWithExt := fmt.Sprintf("%s/%s.chat.txt", io.outputFolder, fileName)
	file, _ := os.Create(fileWithExt)
	_, err := file.Write([]byte(fileContent))
	if err != nil {
		return err
	}

	return nil
}
