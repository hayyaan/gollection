package gollection

type (
	Config struct {
		AppConfig
		DBConfig
	}
	AppConfig struct {
		Name  string
		Usage string
		Host  string
		Port  int
		Debug bool
	}
	DBConfig struct {
		Dialect  string
		Host     string
		Port     int
		Database string
		Username string
		Password string
	}
)
