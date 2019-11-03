package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency takes in a slice of strings and counts the frequency
// of each rune in the total corpus, launching a goroutine to count each
// string concurrently.
func ConcurrentFrequency(input []string) FreqMap {
	freq := FreqMap{}
	channels := make([]chan FreqMap, 0)

	// count each string in parallel, launch a goroutine for each
	for _, piece := range input {
		lineChan := make(chan FreqMap)
		channels = append(channels, lineChan)
		go channeler(piece, lineChan)
	}

	// collect the data from each channel and add it to our
	// frequency map
	for _, channel := range channels {
		data := <-channel
		for key, value := range data {
			if curr, ok := freq[key]; ok {
				freq[key] = value + curr
			} else {
				freq[key] = value
			}
		}
	}

	return freq
}

func channeler(piece string, mapChan chan<- FreqMap) {
	mapChan <- Frequency(piece)
}
