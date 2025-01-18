package utils

import (
	"fmt"
	"os"
)

type FileLogger struct {
	path string
	file *os.File
}

func NewFileLogger(path string) *FileLogger {
	return &FileLogger{
		path: path,
	}
}

func (fl *FileLogger) Writef(format string, a ...any) {
	msg := fmt.Sprintf(format, a...)
	_, err := fl.GetFile().Write([]byte(msg))
	if err != nil {
		panic(err)
	}
}

func (fl *FileLogger) GetFile() *os.File {
	if fl.file != nil {
		return fl.file
	}
	file, err := os.Create(fl.path)
	if err != nil {
		panic(err)
	}
	fl.file = file
	return fl.file
}
