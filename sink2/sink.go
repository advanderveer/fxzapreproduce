package sink2

import (
	"fmt"

	"go.uber.org/fx"
)

// Module exposes all components as dependencies
var Module = fx.Module("redsink",
	fx.Decorate(func(os ...Option) []Option {
		// note this decorator doesn't get called when fx.WithLogger is dependenant on the zap logger
		fmt.Println("DECORATER CALLED!")
		return os
	}),
	fx.Provide(fx.Annotate(New)),
)

type Sink struct{ opts Options }

func New(os ...Option) (s *Sink) {
	s = &Sink{}
	for _, o := range os {
		o(&s.opts)
	}
	fmt.Println("options", s.opts)
	return
}
func (s Sink) Write(p []byte) (n int, err error) { return }
