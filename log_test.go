// Copyright 2023 Emory.Du <orangeduxiaocheng@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package log

import (
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func Test_WithName(t *testing.T) {
	//defer Flush()

	logger := WithName("test")
	logger.Infow("Hello, World!", "foo", "bar")
	logger.Debug("Not Displayed")

	logger = New(&Options{
		OutputPaths:       []string{"stdout"},
		ErrorOutputPaths:  []string{"stderr"},
		Level:             "info",
		Format:            "console",
		DisableCaller:     false,
		DisableStacktrace: false,
		EnableColor:       true,
		Name:              "logger",
	})
	defer logger.Flush()
	logger.Infow("Hello, World!", "foo", "bar")
	logger.Debugw("Hello, World!", "foo", "bar")
	logger.Warnw("Hello, World!", "foo", "bar")
}

func Test_WithValues(t *testing.T) {
	defer Flush()

	logger := WithValues("key", "value") // used for record context
	logger.Info("Hello world!")
	logger.Info("Hello world!")
}

func Test_V(t *testing.T) {
	defer log.Flags()

	V(0).Infow("Hello, World!", "key", "value")
	V(2).Infow("Hello, World!", "key", "value")
}

func Test_Option(t *testing.T) {
	fs := pflag.NewFlagSet("test", pflag.ExitOnError)

	opt := NewOptions()
	opt.AddFlags(fs)

	args := []string{"--log.level=debug"}
	err := fs.Parse(args)
	assert.Nil(t, err)

	assert.Equal(t, "debug", opt.Level)

}
