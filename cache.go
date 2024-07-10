package silo

import "sync"

type Cache struct {
	store  map[string]any
	parent FormatLayer
	once   sync.Once
}

func NewCache(parent FormatLayer) KeyValueLayer {
	return &Cache{
		store:  make(map[string]any),
		parent: parent,
	}
}

func (c *Cache) Get(key string) (any, error) {
	if err := c.load(); err != nil {
		return nil, err
	}

	value, ok := c.store[key]
	if !ok {
		return nil, ErrKeyNotExist
	}
	return value, nil
}

func (c *Cache) Set(key string, value any) error {
	if err := c.load(); err != nil {
		return err
	}

	c.store[key] = value
	return c.parent.Write(c.store)
}

func (c *Cache) Delete(key string) error {
	if err := c.load(); err != nil {
		return err
	}

	delete(c.store, key)
	return c.parent.Write(c.store)
}

func (c *Cache) load() error {
	var err error
	c.once.Do(func() {
		c.store, err = c.parent.Read()
	})
	return err
}
