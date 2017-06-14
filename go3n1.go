package go3n1

import (
	"runtime"
	"sync"
)

func Cnt3n1(n uint64) (cnt uint64) {
	if n == 0 {
		return 0
	}

	for {
		cnt += 1
		if n == 1 {
			break
		}
		if n%2 == 0 {
			n = n >> 1
		} else {
			n = 3*n + 1
		}
	}

	return
}

func List3n1(n uint64) chan uint64 {
	ret := make(chan uint64, 10)
	if n == 0 {
		close(ret)
		return ret
	}
	go func(ch chan uint64) {
		for {
			ch <- n
			if n == 1 {
				break
			}
			if n%2 == 0 {
				n = n >> 1
			} else {
				n = 3*n + 1
			}
		}
		close(ch)
	}(ret)

	return ret
}

type maxlen_ret struct {
	K      uint64
	MaxLen uint64
}

func Maxlen3n1(m, n uint64) (k, maxlen uint64) {
	if m > n {
		m, n = n, m
	}
	goroutines := runtime.NumCPU()
	var wg sync.WaitGroup
	ret := make(chan maxlen_ret, goroutines + 1)

	start := m
	step := (n - m) / uint64(goroutines)
	if step <= 0 {
		step = 1
	}
	for {
		end := start + step
		if end > n {
			end = n
		}
		wg.Add(1)
		go func(sub_m, sub_n uint64) {
			defer wg.Done()

			var x, num, max uint64
			for i := sub_m; i <= sub_n; i++ {
				x = Cnt3n1(i)
				if max == 0 || x > max {
					max = x
					num = i
				}
			}
			if max > 0 {
				ret <- maxlen_ret{
					K:      num,
					MaxLen: max,
				}
			}
		}(start, end)

		start = end
		if start >= n {
			break
		}
	}
	wg.Wait()
	close(ret)
	for v := range ret {
		if maxlen == 0 || v.MaxLen > maxlen {
			maxlen = v.MaxLen
			k = v.K
		}
	}
	return
}
