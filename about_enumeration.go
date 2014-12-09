package go_koans

func aboutEnumeration() {
	{
		var concatenated string
		var total int

		strings := []string{"hello", " world", "!"}
		for i, v := range strings {
			total += i
			concatenated += v
		}

		assert(concatenated == __string__) // for loops have a modern variation
		assert(total == __int__)           // which offers both a value and an index
	}

	{
		var totalLength int

		strings := []string{"hello", " world", "!"}
		for _, v := range strings {
			totalLength += len(v)
		}

		assert(totalLength == __int__) // although we may omit either value
	}
}
