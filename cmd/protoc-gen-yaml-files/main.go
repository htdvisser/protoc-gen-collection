// Copyright 2021 Hylke Visser
// SPDX-License-Identifier: Apache-2.0

package main

import (
	pgs "github.com/lyft/protoc-gen-star"
	"htdvisser.dev/protoc-gen-collection/internal/gendatafiles"
)

func main() {
	pgs.Init(
		pgs.DebugEnv("DEBUG"),
	).RegisterModule(
		gendatafiles.DataFiles(gendatafiles.YAMLEncoder{}),
	).Render()
}
