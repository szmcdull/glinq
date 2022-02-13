package garray

// sortable a generics slice which implements sort.Interface
// inspired by https://github.com/amtoaer/generic-sort
type (
	sortable[T Number]           []T
	sortableDescending[T Number] struct {
		sortable[T]
	}
)

func Sortable[T Number](l []T) sortable[T] {
	return sortable[T](l)
}

func SortableDescending[T Number](l []T) sortableDescending[T] {
	return sortableDescending[T]{l}
}

func (s sortable[T]) Len() int {
	return len(s)
}

func (s sortable[T]) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortable[T]) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortableDescending[T]) Less(i, j int) bool {
	return s.sortable[i] > s.sortable[j]
}
