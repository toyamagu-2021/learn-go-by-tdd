package iteration

func Repeat(character string, iteration int) string {
	var repeated string
	for i := 0; i < iteration; i++ {
		repeated += character
	}
	return repeated
}
