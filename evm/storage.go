package evm

type Storage struct {
	store map[[32]byte][]byte
}

func NewStorage() *Storage {
	return &Storage{
		store: make(map[[32]byte][]byte),
	}
}

func (s *Storage) Get(key [32]byte) []byte {
	return s.store[key]
}

func (s *Storage) Set(key [32]byte, value []byte) {
	s.store[key] = value
}
