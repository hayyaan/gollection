package gollection

type (
	Config struct {
		AppConfig
	}
	AppConfig struct {
		Name  string
		Usage string
		Host  string
		Port  int
	}
)
