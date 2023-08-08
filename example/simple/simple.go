// Copyright 2023 Emory.Du <orangeduxiaocheng@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import "github.com/emorydu/log"

func main() {
	log.Infof("this is a test log, message: %s", "good")
	log.Info("hello, world!")
}
