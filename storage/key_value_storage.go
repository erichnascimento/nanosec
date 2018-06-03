package storage

func NewKeyValueStorage() Storage {
	return &kvStorage{make(map[string]interface{}, 0)}
}

type kvStorage struct {
	documents map[string]interface{}
}

func (s *kvStorage) Set(k string, v interface{}) error {
	s.documents[k] = v
	return nil
}
func (s *kvStorage) Get(k string) (interface{}, error) {
	if d, ok := s.documents[k]; ok {
		return d, nil
	}

	return nil, NotFoundError
}

func (s *kvStorage) Delete(k string) error {
	if _, ok := s.documents[k]; ok {
		delete(s.documents, k)
		return nil
	}

	return NotFoundError
}
