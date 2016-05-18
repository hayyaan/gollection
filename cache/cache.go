package cache

import "time"

type Cache interface {
	Set(k string, v interface{}, d time.Duration) error
	Get(k string) (interface{}, bool)
	Delete(k string) error
	Len() int
}
