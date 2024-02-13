package iteration

func Repeat(character string, repeatNum int) string {
	var repeated string

	for i := 0; i < repeatNum; i++ {
		repeated += character
	}

	return repeated 
}
