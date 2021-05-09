package stats

// Statistics is an interface for fizz-buzz requests' statistics data sources.
type Statistics interface {
	AddEntry(entry Entry) error
	GetMostFrequentEntry() (MostFrequentEntry, error)
}

// Entry gathers the parameters of a fizz-buzz request.
type Entry struct {
	Int1  string `json:"int1"`
	Int2  string `json:"int2"`
	Limit string `json:"limit"`
	Str1  string `json:"str1"`
	Str2  string `json:"str2"`
}

// MostFrequentEntry represents the most frequent fizz-buzz request. It is a
// composition of an Entry and count, i.e., the number of occurrences of this
// entry.
type MostFrequentEntry struct {
	Count int `json:"count"`
	Entry
}
