package camel

func CountLettersInCamelCase(input string) int {
	if len(input) == 0 {
		return 0
	}
	res := 1
	for _, char := range input {
		if 'A' <= char && char <= 'Z' {
			res += 1
		}
	}
	return res
}
