package sink2

import (
	"fmt"
	"os"

	"go.uber.org/fx"
)

// Module exposes all components as dependencies
var Module = fx.Module("sink",
	fx.Decorate(func(os ...Option) []Option {
		// note this decorator doesn't get called when fx.WithLogger is dependenant on the zap logger
		fmt.Println("====== Decorated Called =======")
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
	return
}
func (s Sink) Write(p []byte) (n int, err error) {
	// note: in the reall application this persists to redis
	fmt.Fprintf(os.Stderr, string(p))
	return
}
