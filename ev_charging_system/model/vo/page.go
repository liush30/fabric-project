package vo

type Page[T any] struct {
	Data  []T
	Count int64
}
