package letter

import "sync"

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(text string) FreqMap {
	frequencies := FreqMap{}
	for _, r := range text {
		frequencies[r]++
	}
	return frequencies
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(texts []string) FreqMap {
	var wg sync.WaitGroup
	c := make(chan FreqMap)
	wg.Add(len(texts))

	go func() {
		wg.Wait()
		close(c)
	}()

	for _, t := range texts {
		go func(text string) {
			defer wg.Done()
			c <- Frequency(text)
		}(t)

	}

	res := make(FreqMap)
	for f := range c {
		for k, v := range f {
			res[k] += v
		}
	}
	return res
}
