package gollection

type GollectionEnv interface {
	GetEnv() string
	GetHostPort() (string, int)
}
