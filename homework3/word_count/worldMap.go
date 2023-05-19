package wordcount

func WordCount(strs []string) map[string] int {
	result := map[string] int{}
	for _, str := range strs {
		result[str]++
	}

	return result
}