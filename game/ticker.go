package game

import "time"

func ticker(f func(), interval time.Duration, done <-chan bool) *time.Ticker {
	t := time.NewTicker(interval)
	go func() {
		for {
			select {
			case <-t.C:
				f()
			case <-done:
				return
			}
		}
	}()
	return t
}

func ticker2(f func(...interface{}), args []interface{}, interval time.Duration, done <-chan bool) *time.Ticker {
	t := time.NewTicker(interval)
	go func() {
		for {
			select {
			case <-t.C:
				f(args...)
			case <-done:
				return
			}
		}
	}()
	return t
}
