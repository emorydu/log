package main

import (
	"github.com/emorydu/log"
	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	opts := &log.Options{
		OutputPaths:      []string{"cutter.log"},
		ErrorOutputPaths: []string{"error.log"},
		Level:            "debug",
		Format:           "console",
		Name:             "test",
	}
	opts.Cutter = &lumberjack.Logger{
		Filename:   opts.OutputPaths[0],
		MaxSize:    1, // 单位M
		MaxAge:     3,
		MaxBackups: 30,
		Compress:   false, // 是否压缩
	}
	l := log.New(opts)
	for i := 0; i < 10000; i++ {
		l.Info("Info message")
		l.Debug("Debug message")
		l.Error("Error message")
	}

}
