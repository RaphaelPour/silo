package silo

type Direct struct {
	store map[string]any
}

func NewDirect() Driver {
	return Direct{
		store: make(map[string]any),
	}
}

func (d Direct) Get(key string) (any, error) {
	value, ok := d.store[key]
	if !ok {
		return nil, ErrKeyNotExist
	}
	return value, nil
}

func (d Direct) Set(key string, value any) error {
	d.store[key] = value
	return nil
}

func (d Direct) Delete(key string) error {
	delete(d.store, key)
	return nil
}
