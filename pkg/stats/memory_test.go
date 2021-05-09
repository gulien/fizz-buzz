package stats

import (
	"reflect"
	"testing"
)

func TestInMemoryStatistics_AddEntry(t *testing.T) {
	for i, tc := range []struct {
		entries     []Entry
		expectTotal int
	}{
		{
			entries: []Entry{
				{},
			},
			expectTotal: 1,
		},
		{
			entries: []Entry{
				{},
				{},
			},
			expectTotal: 1,
		},
		{
			entries: []Entry{
				{},
				{},
				{
					Int1: "foo",
				},
			},
			expectTotal: 2,
		},
		{
			entries: []Entry{
				{},
				{},
				{
					Int1: "foo",
				},
				{
					Int1: "foo",
				},
			},
			expectTotal: 2,
		},
		{
			entries: []Entry{
				{},
				{},
				{
					Int1: "foo",
				},
				{
					Int1: "foo",
				},
				{
					Int1: "foo",
				},
			},
			expectTotal: 2,
		},
	} {
		stats := &InMemoryStatistics{
			counter: make(map[Entry]int),
		}

		for _, entry := range tc.entries {
			err := stats.AddEntry(entry)

			if err != nil {
				t.Fatalf("test %d: expected no error but got: %v", i, err)
			}
		}

		actualTotal := len(stats.counter)
		if tc.expectTotal != actualTotal {
			t.Errorf("test %d: expected a total of %d but got %d", i, tc.expectTotal, actualTotal)
		}
	}
}

func TestInMemoryStatistics_GetMostFrequentEntry(t *testing.T) {
	for i, tc := range []struct {
		entries            []Entry
		expectMostFrequent MostFrequentEntry
	}{
		{
			entries: []Entry{},
			expectMostFrequent: MostFrequentEntry{
				Count: 0,
				Entry: Entry{},
			},
		},
		{
			entries: []Entry{
				{},
			},
			expectMostFrequent: MostFrequentEntry{
				Count: 1,
				Entry: Entry{},
			},
		},
	} {
		stats := NewInMemory()

		for _, entry := range tc.entries {
			err := stats.AddEntry(entry)

			if err != nil {
				t.Fatalf("test %d: expected no error but got: %v", i, err)
			}
		}

		actualMostFrequent, err := stats.GetMostFrequentEntry()
		if err != nil {
			t.Fatalf("test %d: expected no error but got: %v", i, err)
		}

		if !reflect.DeepEqual(tc.expectMostFrequent, actualMostFrequent) {
			t.Errorf("test %d: expected most frequent entry %v, but got: %v", i, tc.expectMostFrequent, actualMostFrequent)
		}
	}
}
