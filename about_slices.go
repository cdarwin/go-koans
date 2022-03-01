package go_koans

func aboutSlices() {
	fruits := []string{"apple", "orange", "mango"}

	assert(fruits[0] == __string__) // slices seem like arrays
	assert(len(fruits) == __int__)  // in nearly all respects

	tasty_fruits := fruits[1:3]           // we can even slice slices
	assert(tasty_fruits[0] == __string__) // slices of slices also share the underlying data

	hero_slots := []string{"Leia", "Luke", "Ray"}
	assert(cap(hero_slots) == __int__) // the capacity is initially the length

	hero_slots = append(hero_slots, "Han")
	assert(len(hero_slots) == __int__) // slices can be extended with append(), much like realloc in C
	assert(cap(hero_slots) == __int__) // but with better optimizations

	hero_slots = append(hero_slots, "We don't", "need another", "hero")

	assert(len(hero_slots) == __int__) // append() can take N arguments to append to the slice
	assert(cap(hero_slots) == __int__) // the capacity optimizations have a guessable algorithm
}
