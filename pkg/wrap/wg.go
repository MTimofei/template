package wrap

import "sync"

type W struct {
	wg    *sync.WaitGroup
	errCh chan error
}

func NewW() *W {
	return &W{
		wg:    new(sync.WaitGroup),
		errCh: make(chan error),
	}
}

func (w *W) Run(nex func() error) {
	w.wg.Add(1)
	defer w.wg.Done()
	nex()
}
