package storage

type inMemoryDatabase interface {
	Set(key any, value any)
	Get(key any) any
	Delete(key any)
	Start_transaction()
	Commit()
	Rollback()
}

type data struct {
	records map[any]any
	parent  *data
	child   *data
}

type Storage struct {
	datastore data
	current   *data
}

func (s *Storage) dropLast() {
	if s.isRoot() {
		return // Root point
	}
	s.current = s.current.parent
	s.current.child.parent = nil
	s.current.child = nil
}

func (s *Storage) isRoot() bool {
	return s.current.parent == nil
}

func (s *Storage) Start_transaction() {
	newData := &data{
		records: map[any]any{},
		parent:  s.current,
	}
	s.current.child = newData
	s.current = newData
}

func (s *Storage) Commit() {
	if s.isRoot() {
		return
	}

	// Merge to parent
	for key, val := range s.current.records {
		s.current.parent.records[key] = val
	}

	s.dropLast()
}

func (s *Storage) Rollback() {
	s.dropLast()
}

func (s *Storage) Delete(key any) {
	s.current.records[key] = nil
}

func (s *Storage) Set(key any, value any) {
	s.current.records[key] = value
}

func (s *Storage) Get(key any) any {
	current := s.current

	for {
		val, ok := current.records[key]

		if ok || current.parent == nil {
			return val
		}

		current = current.parent
	}
}

func InMemoryDatabase() inMemoryDatabase {
	data := data{
		records: map[any]any{},
	}

	return &Storage{
		datastore: data,
		current:   &data,
	}
}
