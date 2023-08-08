// Copyright 2023 Emory.Du <orangeduxiaocheng@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"flag"
	"github.com/emorydu/log"
	"time"
)

var (
	h bool

	level  int
	format string
)

func main() {
	flag.BoolVar(&h, "h", false, "Print this help.")
	flag.IntVar(&level, "l", 0, "Log level.")
	flag.StringVar(&format, "f", "console", "Log output format.")

	flag.Parse()

	if h {
		flag.Usage()

		return
	}

	// logger configuration
	opts := &log.Options{
		Level:            "debug",
		Format:           "console",
		EnableColor:      true, // if you need output to local path, with EnableColor must be false.
		DisableCaller:    true,
		OutputPaths:      []string{"test.log", "stdout"},
		ErrorOutputPaths: []string{"error.log"},
		Cutter:           nil,
	}
	/*
		tips: Use the log split by size feature, uncomment this comment

		opts.Cutter = &lumberjack.Logger{
			Filename:   opts.OutputPaths[0],
			MaxSize:    1,
			MaxAge:     3,
			MaxBackups: 30,
			Compress:   false,
		}
	*/

	// Initialize the global logger
	log.Init(opts)
	defer log.Flush()

	time.Sleep(5 * time.Second)
	// Debug, Info(with filed), Warnf, Errorw
	log.Debug("This is a debug message")
	log.Info("This is info message", log.Int32("int_key", 10))
	log.Warnf("This is a formatted %s message", "warn")
	log.Errorw("Message printed with Errorw", "X-Request-ID", "fff54504-64da-4088-9b86-67824a7fb508")

	// WithValues
	lv := log.WithValues("X-Request-ID", "fbf54504-64da-4088-9b86-67824a7fb508")
	lv.Infow("Info message printed with [WithValues] logger")
	lv.Debugw("Debug message printed with [WithValues] logger")

	// Context
	ctx := lv.WithContext(context.Background())
	lc := log.FromContext(ctx)
	lc.Info("Message printed with [WithContext] logger")

	// Name
	ln := lv.WithName("test")
	ln.Info("Message printed with [WithName] logger")

	// V level
	log.V(log.InfoLevel).Info("This is a V level message")
	log.V(log.ErrorLevel).Infow(
		"This is a V level message with fields", "X-Request-ID", "9bf54504-64da-4088-9b86-67824a7fb508")

}
