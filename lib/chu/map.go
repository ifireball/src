package chu

func Map[inT, outT interface{}](ch <-chan inT, f func(inT) (outT, bool)) <-chan outT {
	out := make(chan outT)
	collector := make(chan struct {
		out  outT
		keep bool
	})
	go func() {
		defer func() { close(out) }()
		defer func() { close(collector) }()
		count := 0
		for o := range ch {
			go func(o inT) {
				oo, keep := f(o)
				collector <- struct {
					out  outT
					keep bool
				}{oo, keep}
			}(o)
			count++
		}
		for i := 0; i < count; i++ {
			cs := <-collector
			if cs.keep {
				out <- cs.out
			}
		}
	}()
	return out
}
