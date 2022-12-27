package iteration

func Repeat(s string, times int) string {
	result := ""
	
	for i := 0; i < times; i++ {
		result += s
	}

	return result
}