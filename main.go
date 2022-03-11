package main

import (
	"os"

	"github.com/advanderveer/fxzapreproduce/log2"
	"github.com/advanderveer/fxzapreproduce/sink2"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

// app is our top level module
var app = fx.Module("app",
	sink2.Module,
	log2.Module,
	fx.Invoke(func(_ *zap.Logger) {}),
)

// main runs the app
func main() {
	fx.New(FxLogger(), app).Run()
}

// FxLogger is a convenient option that configures fx to use the zap logger.
func FxLogger() fx.Option {
	return fx.WithLogger(func(l *zap.Logger) fxevent.Logger {
		return &fxevent.ConsoleLogger{W: os.Stderr}
	})
}
