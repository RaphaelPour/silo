package silo

type Cache struct {
	parent FormatLayer
	m      map[string]any
}

func NewCache(parent FormatLayer) KeyValueLayer {
	return Cache{
		parent: parent,
	}
}

func (c Cache) Get(key string) (any, error) {
	var err error
	if c.m == nil {
		c.m, err = c.parent.Read()
		if err != nil {
			return nil, err
		}
	}

	return c.m[key], nil
}

func (c Cache) Set(key string, value any) error {
	if c.m == nil {
		var err error
		c.m, err = c.parent.Read()
		if err != nil {
			return err
		}
	}
	c.m[key] = value
	return c.parent.Write(c.m)
}

func (c Cache) Delete(key string) error {
	delete(c.m, key)
	return c.parent.Write(c.m)
}
