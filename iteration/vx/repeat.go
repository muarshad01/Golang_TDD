package iteration

const repeatCount = 5

// Repeat returns character repeated 5 times.
func Repeat(str string) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += str
	}
	return repeated
}
