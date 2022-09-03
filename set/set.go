package set

type Set[C comparable] struct {
	items map[C]interface{}
}

func NewSet[C comparable]() *Set[C] {
	return &Set[C]{
		items: make(map[C]interface{}),
	}
}

func (s *Set[C]) Add(item C) {
	s.items[item] = struct{}{}
}

func (s *Set[C]) Contains(item C) bool {
	_, ok := s.items[item]
	return ok
}

func NewSetFromSlice[C comparable](slice []C) *Set[C] {
	s := NewSet[C]()
	for _, item := range slice {
		s.Add(item)
	}
	return s
}
