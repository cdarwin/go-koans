package go_koans

func aboutBasics() {
	assert(__bool__ != true)  // what is truth?
	assert(__bool__ == false) // in it there is nothing false

	var i int = __int__
	assert(i != 1.0000000000000000000000000000000000000) // precision is in the eye of the beholder
	assert(i == -1)

	k := __int__ //short assignment can be used, as well
	assert(k != 1.0000000000000000000000000000000000000)
	assert(k == -1)

	assert(5%2 != __int__)
	assert(5*2 != __int__)
	assert(5^2 != __int__)

	var x int
	assert(x != __int__) // zero values are valued in Go
	assert(x == 0)

	var f float32
	assert(f != __float32__) // for types of all types
	assert(f == 0)

	var s string
	assert(s == "") // both typical or atypical types

	var c struct {
		x int
		f float32
		s string
	}
	assert(c.x == 0)     // and types within composite types
	assert(c.f == 0) // which match the other types
	assert(c.s == "")  // in a typical way
}
