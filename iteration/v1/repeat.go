package iteration

// Repeat returns character repeated 5 times.
func Repeat(str string) string {
	var repeated string
	for i := 0; i < 5; i++ {
		repeated += str
	}
	return repeated
}
