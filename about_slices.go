package gokoans

func aboutSlices() {
	fruits := []string{"apple", "orange", "mango"}

	assert(fruits[0] == __string__) // slices seem like arrays
	assert(len(fruits) == __int__)  // in nearly all respects

	tastyFruits := fruits[1:3]           // we can even slice slices
	assert(tastyFruits[0] == __string__) // slices of slices also share the underlying data

	pregnancySlots := []string{"baby", "baby", "lemon"}
	assert(cap(pregnancySlots) == __int__) // the capacity is initially the length

	pregnancySlots = append(pregnancySlots, "baby!")
	assert(len(pregnancySlots) == __int__) // slices can be extended with append(), much like realloc in C
	assert(cap(pregnancySlots) == __int__) // but with better optimizations

	pregnancySlots = append(pregnancySlots, "another baby!?", "yet another, oh dear!", "they must be Catholic")

	assert(len(pregnancySlots) == __int__) // append() can take N arguments to append to the slice
	assert(cap(pregnancySlots) == __int__) // the capacity optimizations have a guessable algorithm
}
