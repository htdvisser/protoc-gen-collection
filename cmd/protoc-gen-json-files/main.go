package main

import (
	pgs "github.com/lyft/protoc-gen-star"
	"htdvisser.dev/protoc-gen-collection/internal/gendatafiles"
)

func main() {
	pgs.Init(
		pgs.DebugEnv("DEBUG"),
	).RegisterModule(
		gendatafiles.DataFiles(gendatafiles.JSONEncoder{}),
	).Render()
}
