package trees

type node[T comparable] struct {
	val T

	right *node[T]
	left  *node[T]
}

type Tree[T comparable] interface {
	Insert(item T)

	Delete(item T) error

	Search(item T) (*T, error)

	InOrder() []T
}
