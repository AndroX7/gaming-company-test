package middleware

type Config struct {
	// tbh i want to add jwt config, redis cache, and some grouping route here but it looks like the requirement doesn't need it so i don't create it
	// then let just it be empty struct
}

type Middleware struct {
	config Config
}

func New(cfg Config) *Middleware {
	return &Middleware{
		config: cfg,
	}
}
