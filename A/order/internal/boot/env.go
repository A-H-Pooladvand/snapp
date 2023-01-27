package boot

import "github.com/joho/godotenv"

type Env struct{}

// Boot is an adapter method to bootstrap
func (e *Env) Boot() error {
	return e.Load()
}

// Load loads the .env file if it exists.
func (e *Env) Load() error {
	return godotenv.Load(".env")
}
