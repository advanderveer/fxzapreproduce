package main

import (
	"context"
	"testing"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

func TestDecorateWithoutFxLogger(t *testing.T) {
	app := fx.New(app, fx.Invoke(func(_ *zap.Logger) {}))
	ctx := context.Background()
	if err := app.Start(ctx); err != nil {
		panic(err)
	}
}

func TestDecorateWithFxLogger(t *testing.T) {
	app := fx.New(FxLogger(), app, fx.Invoke(func(_ *zap.Logger) {}))
	ctx := context.Background()
	if err := app.Start(ctx); err != nil {
		panic(err)
	}
}
