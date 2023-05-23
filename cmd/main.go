package main

import (
	"github.com/tachunwu/natsfx"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		natsfx.Module,
	)

	app.Run()
}
