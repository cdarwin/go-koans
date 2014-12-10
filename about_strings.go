package go_koans

import "fmt"

func aboutStrings() {
	assert("a"+__string__ == "abc") // string concatenation need not be difficult
	assert(len("abc") == __int__)   // and bounds are thoroughly checked

	assert("abc"[0] == __byte__) // their contents are reminiscent of C

	assert("smith"[2:] == __string__)  // slicing may omit the end point
	assert("smith"[:4] == __string__)  // or the beginning
	assert("smith"[2:4] == __string__) // or neither
	assert("smith"[:] == __string__)   // or both

	assert("smith" == __string__) // they can be compared directly
	assert("smith" < __string__)  // i suppose maybe this could be useful.. someday

	bytes := []byte{'a', 'b', 'c'}
	assert(string(bytes) == __string__) // strings can be created from byte-slices

	bytes[0] = 'z'
	assert(string(bytes) == __string__) // byte-slices can be mutated, although strings cannot

	assert(fmt.Sprintf("hello %s", __string__) == "hello world") // our old friend sprintf returns
	assert(fmt.Sprintf("hello \"%s\"", "world") == __string__)   // quoting is familiar
	assert(fmt.Sprintf("hello %q", "world") == __string__)       // although it can be done more easily

	assert(fmt.Sprintf("your balance: %d and %0.2f", 3, 4.5589) == __string__) // "the root of all evil" is actually a misquotation, by the way
}
