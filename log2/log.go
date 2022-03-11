package log2

import (
	"github.com/advanderveer/fxzapreproduce/sink2"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Module = fx.Module("log",
	fx.Provide(
		zap.New,
		zapcore.NewCore,
		zap.NewProductionEncoderConfig,
		zapcore.NewJSONEncoder,
		fx.Annotate(NewEnabler),
		fx.Annotate(WriteSyncers),
	),
)

func NewEnabler() zapcore.LevelEnabler {
	return zap.DebugLevel
}

func WriteSyncers(rs *sink2.Sink) zapcore.WriteSyncer {
	return zapcore.AddSync(rs)
}
