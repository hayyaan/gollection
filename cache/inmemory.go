package cache

import (
	"time"

	gocache "github.com/patrickmn/go-cache"
)

type InMemory struct {
	gocache *gocache.Cache
}

func NewInMemory() *InMemory {
	return &InMemory{
		gocache: gocache.New(5*time.Minute, 30*time.Second),
	}
}

func (c *InMemory) Set(k string, v interface{}, d time.Duration) error {
	c.gocache.Set(k, v, d)
	return nil
}

func (c *InMemory) Get(k string) (interface{}, bool) {
	return c.gocache.Get(k)
}

func (c *InMemory) Delete(k string) error {
	c.gocache.Delete(k)
	return nil
}

func (c *InMemory) Len() int {
	return c.gocache.ItemCount()
}
