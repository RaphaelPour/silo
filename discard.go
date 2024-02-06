package silo

type Discard struct{}

func NewDiscard() KeyValueLayer {
	return Discard{}
}

func (d Discard) Get(key string) (any, error) {
	return nil, nil
}

func (d Discard) Set(key string, value any) error {
	return nil
}

func (d Discard) Delete(key string) error {
	return nil
}
