package silo

import (
	"errors"
	"fmt"
	"os"
	"time"
)

var (
	ErrOpenLogFile = errors.New("error opening log file")
)

func prettifyErr(err error) string {
	if err == nil {
		return "<nil>"
	}
	return err.Error()
}

type Log struct {
	parent Driver
	file   *os.File
}

func NewLog(parent Driver, filename string) (Driver, error) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return nil, fmt.Errorf("%w %s: %w", ErrOpenLogFile, filename, err)
	}

	return Log{
		parent: parent,
		file:   file,
	}, nil
}

func (l Log) write(msg string) {
	l.file.WriteString(fmt.Sprintf("%s %s\n", time.Now().Format(time.DateTime), msg)) //nolint: errcheck
}

func (l Log) Get(key string) (any, error) {
	value, err := l.parent.Get(key)
	l.write(fmt.Sprintf("[   GET] %s=%v err=%s", key, value, prettifyErr(err)))
	return value, err
}

func (l Log) Set(key string, value any) error {
	err := l.parent.Set(key, value)
	l.write(fmt.Sprintf("[   SET] %s=%v err=%s", key, value, prettifyErr(err)))
	return err
}

func (l Log) Delete(key string) error {
	err := l.parent.Delete(key)
	l.write(fmt.Sprintf("[DELETE] %s err=%s", key, prettifyErr(err)))
	return err
}
