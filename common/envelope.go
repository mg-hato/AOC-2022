package common

type Envelope[T any] interface {
	Get() T
}
