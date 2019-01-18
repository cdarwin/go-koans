package go_koans

func aboutBasics() {
	assert(__bool__ == true)  // what is truth?
	assert(__bool__ != false) // in it there is nothing false

	var i int = __int__
	assert(i == 1.0000000000000000000000000000000000000) // precision is in the eye of the beholder

	k := __int__ //short assignment can be used, as well
	assert(k == 1.0000000000000000000000000000000000000)

	assert(5%2 == __int__)
	assert(5*2 == __int__)
	assert(5^2 == __int__)

	var x int
	assert(x == __int__) // listen to the darkness of an unset variable

	var f float32
	assert(f == __float32__) // what is the code that is not written?

	var s string
	assert(s == __string__) // consider the emptiness of a string

	var c struct {
		x int
		f float32
		s string
	}
	assert(c.x == __int__)     // create meaning from emptiness
	assert(c.f == __float32__) // undefined structure isn't
	assert(c.s == __string__)  // what is the structure of a string?
}
