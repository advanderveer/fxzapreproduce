package sink2

// Options configure Foo
type Options struct {
	A string
	B int
}

// WithA adds an option
func WithA(v string) Option {
	return func(o *Options) { o.A = v }
}

// WithB adds an option
func WithB(v int) Option {
	return func(o *Options) { o.B = v }
}

// Option type follows the functional option pattern
type Option func(o *Options)
