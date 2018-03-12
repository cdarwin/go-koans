package go_koans

import "fmt"

func aboutArrays() {
	fruits := [4]string{"apple", "orange", "mango"}

	assert(fruits[0] == __string__) // indexes begin at 0
	assert(fruits[1] == __string__) // one is indeed the loneliest number
	assert(fruits[2] == __string__) // it takes two to ...tango?
	assert(fruits[3] == __string__) // there is no spoon, only an empty value

	assert(len(fruits) == __int__) // the length is what the length is
	assert(cap(fruits) == __int__) // it can hold no more

	assert(fruits == [4]string{}) // comparing arrays is not like comparing apples and oranges

	tasty_fruits := fruits[1:3]                           // defining oneself as a variation of another
	assert(fmt.Sprintf("%T", tasty_fruits) == __string__) //and get not a simple array as a result
	assert(tasty_fruits[0] == __string__)                 // slices of arrays share some data
	assert(tasty_fruits[1] == __string__)                 // albeit slightly askewed

	assert(len(tasty_fruits) == __int__) // its length is manifest
	assert(cap(tasty_fruits) == __int__) // but its capacity is surprising!

	tasty_fruits[0] = "lemon" // are their shared roots truly identical?

	assert(fruits[0] == __string__) // has this element remained the same?
	assert(fruits[1] == __string__) // how about the second?
	assert(fruits[2] == __string__) // surely one of these must have changed
	assert(fruits[3] == __string__) // but who can know these things

	veggies := [...]string{"carrot", "pea"}

	assert(len(veggies) == __int__) // array literals need not repeat an obvious length
}
