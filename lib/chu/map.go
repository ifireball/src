package chu

func Map[inT, outT interface{}](ch <-chan inT, f func(inT) outT) <-chan outT {
	out := make(chan outT)
	go func() {
		defer func() { close(out) }()
		for o := range ch {
			out <- f(o)
		}
	}()
	return out
}
