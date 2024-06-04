package chu

func Map[inT, outT interface{}](ch <-chan inT, f func(inT) (outT, bool)) <-chan outT {
	out := make(chan outT)
	go func() {
		defer func() { close(out) }()
		for o := range ch {
			if oo, keep := f(o); keep {
				out <- oo
			}
		}
	}()
	return out
}
