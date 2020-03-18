package tsync

import "sync"

func ExecuteAll(functions ...func() interface{}) map[int]interface{} {
	m := make(map[int]chan interface{})
	for i, f := range functions {
		ch := make(chan interface{}, 1)
		m[i] = ch
		exf := f
		go func() {
			ch <- exf()
		}()
	}
	rm := make(map[int]interface{})
	for k, ch := range m {
		rm[k] = <-ch
	}
	return rm
}

func SubmitAll(functions ...func()) {
	var wg sync.WaitGroup
	wg.Add(len(functions))
	for _, f := range functions {
		exf := f
		go func() {
			defer wg.Done()
			exf()
		}()
	}
	wg.Wait()
}

func IterateSubmit(iterations int, functions func()) {
	var wg sync.WaitGroup
	wg.Add(iterations)
	for i := 0; i < iterations; i++ {
		go func() {
			functions()
			defer wg.Done()
		}()
	}
	wg.Wait()
}
