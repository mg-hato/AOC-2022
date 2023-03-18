package envelope

type Envelope[T any] interface {
	Get() T
}
