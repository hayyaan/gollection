package gollection

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

// LoadEnv takes a filename, should be .env, and loads its contents,
// which are then set as ENV variables
func LoadEnv(filename string) error {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return nil
	}

	f, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	vars := parseEnvFile(string(f))
	for k, v := range vars {
		if _, exists := os.LookupEnv(k); !exists {
			os.Setenv(k, v)
		}
	}

	return nil
}

// parseEnvFile gets the contents of a .env file and creates a map
func parseEnvFile(e string) map[string]string {
	vars := map[string]string{}

	for _, line := range strings.Split(e, "\n") {
		if len(line) > 0 {
			if v := strings.Split(line, "="); len(v) == 2 {
				vars[v[0]] = v[1]
			}
		}
	}

	return vars
}

// GetEnv gets a ENV value as string or returns the fallback
func GetEnv(key string, fallback string) string {
	s := os.Getenv(key)
	if s != "" {
		return s
	}

	return fallback
}

// GetEnvInt gets a ENV value as int or returns the fallback
func GetEnvInt(key string, fallback int) int {
	s := GetEnv(key, "")
	if s != "" {
		i, err := strconv.Atoi(s)
		if err != nil {
			log.Printf("Couldn't parse env %s, falling back to %d", key, fallback)
			return fallback
		}

		return i
	}

	return fallback
}

// GetEnvBool gets a ENV value as bool or returns the fallback
func GetEnvBool(key string, fallback bool) bool {
	s := GetEnv(key, "")
	if s != "" {
		b, err := strconv.ParseBool(s)
		if err != nil {
			log.Printf("Couldn't parse env %s, falling back to %t", key, fallback)
			return fallback
		}

		return b
	}

	return fallback
}
