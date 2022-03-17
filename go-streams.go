package streams

type FilterFunc[T any] func(v T) bool

type MapFunc[T, R any] func(v T) R

type Stream[T any] struct {
	slice   []T
	filters filters[T]
}

func New[T any](slice []T) *Stream[T] {
	return &Stream[T]{
		slice: slice,
	}
}

func (s *Stream[T]) Filter(f FilterFunc[T]) *Stream[T] {
	s.filters = append(s.filters, f)
	return s
}

func (s *Stream[T]) Collect() []T {
	return s.filters.Collect(s.slice)
}

type MappedStream[T any, W any] struct {
	s *Stream[W]
	f MapFunc[W, T]

	filters filters[T]
}

func Map[T any, W any](s *Stream[W], f MapFunc[W, T]) *MappedStream[T, W] {
	return &MappedStream[T, W]{
		s: s,
		f: f,
	}
}

func (s *MappedStream[T, W]) Filter(f FilterFunc[T]) *MappedStream[T, W] {
	s.filters = append(s.filters, f)
	return s
}

func (s *MappedStream[T, W]) Collect() []T {
	var slice []T
	for _, element := range s.s.Collect() {
		slice = append(slice, s.f(element))
	}

	return s.filters.Collect(slice)
}

type filters[T any] []FilterFunc[T]

func (f filters[T]) Collect(slice []T) []T {
	var result []T
elements:
	for _, element := range slice {
		for _, filter := range f {
			if !filter(element) {
				continue elements
			}
		}
		result = append(result, element)
	}
	return result
}
