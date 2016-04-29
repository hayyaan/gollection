package gollection

type GollectionConfig interface {
	GetEnv() string
	GetAddrPort() (string, int)
}
