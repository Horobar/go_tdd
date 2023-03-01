package iteration

func Repeat(charakter string, times int) string {

	var repeated string
	for i := 0; i < times; i++ {
		repeated += charakter
	}
	return repeated
}
