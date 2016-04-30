package gollection

type GollectionConfig interface {
	GetEnv() string
	GetHostPort() (string, int)
}
