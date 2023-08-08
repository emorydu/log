// Copyright 2023 Emory.Du <orangeduxiaocheng@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package logrus

import (
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"io/ioutil"
)

// NewLogger create a logrus logger, add hook to it and return it.
func NewLogger(zapLogger *zap.Logger) *logrus.Logger {
	logger := logrus.New()
	logger.SetOutput(ioutil.Discard)
	logger.AddHook(newHook(zapLogger))

	return logger
}
