package log2

import (
	"fmt"
	"os"

	"github.com/advanderveer/fxzapreproduce/sink2"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Module = fx.Module("log2",
	// fx.Provide(fx.Annotate(options.ParseEnv[opts])),
	fx.Provide(
		zap.New,
		zapcore.NewCore,
		zap.NewProductionEncoderConfig,
		zapcore.NewJSONEncoder,
		fx.Annotate(NewEnabler),
		fx.Annotate(WriteSyncers, fx.ParamTags(`optional:"true"`)),
	),
)

func NewEnabler() zapcore.LevelEnabler {
	return zap.DebugLevel
}

func WriteSyncers(rs *sink2.Sink) zapcore.WriteSyncer {
	if rs == nil {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stderr))
	}
	fmt.Println("TRIGGER multi writer syncer 2")
	return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stderr), zapcore.AddSync(rs))
}
