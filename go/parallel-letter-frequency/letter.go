package letter

import "sync"

const testVersion = 1

func ConcurrentFrequency(strs []string) FreqMap {
	m := FreqMap{}
	quit := make(chan bool)
	var mutex = &sync.Mutex{}

	for _, str := range strs {
		go func(s string) {
			fm := Frequency(s)
			mutex.Lock()
			for k, v := range fm {
				m[k] += v
			}
			mutex.Unlock()
			quit <- true
		}(str)
	}

	finished := 0

	for finished != len(strs) {
		select {
		case <-quit:
			finished++
		}
	}
	close(quit)
	return m
}
