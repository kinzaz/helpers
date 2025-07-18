package slices

import "sync"

func Each[I any](inp []I, eachFunc func(int, I)) {
	for i := range inp {
		eachFunc(i, inp[i])
	}
}

func TryEach[I any](inp []I, eachFunc func(int, I) error) error {
	var err error
	for i := range inp {
		if err = eachFunc(i, inp[i]); err != nil {
			return err
		}
	}
	return nil
}

// GoEach вызывает eachFunc в многопоточном режиме. Если threadNum больше ноля, количество потоков будет ограничено.
func GoEach[I any](inp []I, threadNum int, eachFunc func(int, I)) {
	var threads chan struct{}
	limited := threadNum > 0
	wg := new(sync.WaitGroup)

	wg.Add(len(inp))

	if limited {
		threads = make(chan struct{}, threadNum)
	}

	Each(inp, func(i int, v I) {
		if limited {
			threads <- struct{}{}
		}

		go func(ii int, vv I) {
			defer func() {
				if limited {
					<-threads
				}
				wg.Done()
			}()

			eachFunc(ii, vv)
		}(i, v)
	})

	wg.Wait()
	if limited {
		close(threads)
	}
}

func GoTryEach[I any](inp []I, threadNum int, eachFunc func(int, I) error) <-chan error {
	var threads chan struct{}
	limited := threadNum > 0

	errs := make(chan error, len(inp))

	go func() {
		defer close(errs)

		wg := new(sync.WaitGroup)

		wg.Add(len(inp))

		if limited {
			threads = make(chan struct{}, threadNum)
		}

		Each(inp, func(i int, v I) {
			if limited {
				threads <- struct{}{}
			}

			go func(ii int, vv I) {
				defer func() {
					if limited {
						<-threads
					}
					wg.Done()
				}()

				if err := eachFunc(ii, vv); err != nil {
					errs <- err
				}
			}(i, v)
		})

		wg.Wait()
		if limited {
			close(threads)
		}
	}()

	return errs
}
