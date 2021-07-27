package letter

type FreqMap map[rune]int

func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func worker(s string, c chan FreqMap) {
	c <- Frequency(s)
}

func ConcurrentFrequency(s []string) FreqMap {
	var c chan FreqMap = make(chan FreqMap, 2)

	for _, v := range s {
		go worker(v, c)
	}

	frequencies := FreqMap{}
	for _ = range s {
		f := <-c
		for k, v := range f {
			if _, ok := frequencies[k]; !ok {
				frequencies[k] = 0
			}
			frequencies[k] += v
		}
	}

	return frequencies
}
