package iteration

// Repeat repeats text specified times
func Repeat(text string, times int) (repeated string) {
	for i := 0; i < times; i++ {
		repeated = repeated + text
	}
	return
}
