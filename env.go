package gollection

// Env makes sure that all env variables are available to the application.
type Env interface {
	// GetEnv returns production, local...
	GetEnv() string

	// GetHostPort returns the http host & port
	GetHostPort() (string, int)
}
