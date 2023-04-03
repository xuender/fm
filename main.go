package main

import (
	"embed"

	"github.com/xuender/fm/cmd"
	"github.com/youthlin/t"
)

//go:embed locales
var _locales embed.FS

func main() {
	t.LoadFS(_locales)
	cmd.Execute()
}
