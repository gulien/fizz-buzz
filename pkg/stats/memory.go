package stats

import "sync"

// InMemoryStatistics is an in memory data source for fizz-buzz requests'
// statistics.
type InMemoryStatistics struct {
	mostFrequent Entry
	counter      map[Entry]int

	mu sync.RWMutex
}

// NewInMemory returns an instance of InMemoryStatistics.
func NewInMemory() Statistics {
	return &InMemoryStatistics{
		counter: make(map[Entry]int),
	}
}

// AddEntry adds a fizz-buzz request to memory.
func (stats *InMemoryStatistics) AddEntry(entry Entry) error {
	stats.mu.Lock()
	defer stats.mu.Unlock()

	mostFrequentCount, ok := stats.counter[stats.mostFrequent]
	if !ok {
		// No most frequent, so this entry is the first ever.
		stats.mostFrequent = entry
		stats.counter[entry] = 1

		return nil
	}

	if stats.mostFrequent == entry {
		// This entry is the same as the most frequent.
		stats.counter[stats.mostFrequent] = mostFrequentCount + 1

		return nil
	}

	count, ok := stats.counter[entry]
	if !ok {
		// First time for this entry.
		stats.counter[entry] = 1

		return nil
	}

	count += 1

	// Older prevail if equal.
	if count <= mostFrequentCount {
		stats.counter[entry] = count

		return nil
	}

	// New most frequent entry.
	stats.mostFrequent = entry
	stats.counter[entry] = count

	return nil
}

// GetMostFrequentEntry returns the most frequent fizz-buzz request from
// memory.
func (stats *InMemoryStatistics) GetMostFrequentEntry() (MostFrequentEntry, error) {
	stats.mu.RLock()
	defer stats.mu.RUnlock()

	countMostFrequent, ok := stats.counter[stats.mostFrequent]

	if !ok {
		// No most frequent entry yet, returns an "empty" entry.
		return MostFrequentEntry{
			Count: 0,
			Entry: stats.mostFrequent, // = Entry{}.
		}, nil
	}

	return MostFrequentEntry{
		Count: countMostFrequent,
		Entry: stats.mostFrequent,
	}, nil
}

// Interface guards.
var (
	_ Statistics = (*InMemoryStatistics)(nil)
)
