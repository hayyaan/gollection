package database

// Config describes how to connect to a database
type Config struct {
	Username string
	Password string
	Host     string
	Port     int
	Database string
}
