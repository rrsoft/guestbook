package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func Write(s string) error {
	filename := logFile()
	return WriteTo(filename, s)
}

func WriteTo(filename, s string) error {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer f.Close()
	std := log.New(f, "", log.LstdFlags)
	return std.Output(2, s)
}

func WriteLog(filename string, data []byte, perm os.FileMode) error {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, perm)
	if err != nil {
		return err
	}
	n, err := f.Write(data)
	f.Close()
	if err == nil && n < len(data) {
		err = io.ErrShortWrite
	}
	return err
}

func logFile() string {
	now := time.Now()
	path := makeDir(now)
	return fmt.Sprintf("%s/%d.log", path, now.Day())
}

func dateString() string {
	return time.Now().Format("2006-01-02")
}

func makeDir(t time.Time) string {
	path := fmt.Sprintf("log/%d/%d", t.Year(), t.Month())
	os.MkdirAll(path, 0755)
	return path
}
