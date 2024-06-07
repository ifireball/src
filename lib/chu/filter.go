package chu

func Filter[T interface{}](ch <-chan T, f func(T) bool) <- chan T {
	return Map(ch, func(o T) (T, bool) {
		return o, f(o)
	})
}
