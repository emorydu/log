// Copyright 2023 Emory.Du <orangeduxiaocheng@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"flag"
	"github.com/emorydu/log"
)

var (
	h bool

	level  int
	format string
)

func main() {
	flag.BoolVar(&h, "h", false, "Print this help.")
	flag.IntVar(&level, "l", 0, "Log level.")
	flag.StringVar(&format, "f", "console", "log output format.")

	flag.Parse()

	if h {
		flag.Usage()

		return
	}

	opts := &log.Options{
		OutputPaths:      []string{"test.log", "stdout"},
		ErrorOutputPaths: []string{"error.log"},
		Level:            "debug",
		Format:           "console",
		EnableColor:      true,
		DisableCaller:    true,
	}

	log.Init(opts)
	defer log.Flush()

	lv := log.WithValues("X-Request-ID", "7a7b9f24-4cae-4b2a-9464-69088b45b904")
	lv.Infof("Start to call printString function.")
	ctx := lv.WithContext(context.Background())
	printString(ctx, "World")
}

func printString(ctx context.Context, s string) {
	lc := log.FromContext(ctx)
	lc.Infof("Hello %s", s)
	lc.Errorf("Error Message %s", s)
}
