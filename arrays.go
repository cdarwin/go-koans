package go_koans

var __string__ string = "impossibly lame value"
var __int__ int = -1

func testArrays() {
  fruits := [4]string{"apple", "orange", "mango"}

  assert(fruits[0] == __string__) // indexes begin at 0
  assert(fruits[1] == __string__) // one is indeed the loneliest number
  assert(fruits[2] == __string__) // it takes two to ...tango?
  assert(fruits[3] == __string__) // there is no spoon

  assert(len(fruits) == __int__) // the length is what the length is
  assert(cap(fruits) == __int__) // it can hold no more

  tasty_fruits := fruits[1:3] // defining ones self as a variation of another
  assert(tasty_fruits[0] == __string__) // slices of arrays share some identity
  assert(tasty_fruits[1] == __string__) // albeit slightly askewed

  tasty_fruits[0] = "lemon" // but are their shared roots truly identical?

  assert(fruits[0] == __string__) // has this remained the same?
  assert(fruits[1] == __string__) // how about this?
  assert(fruits[2] == __string__) // surely one of these must have changed
  assert(fruits[3] == __string__) // but who can know these things
}
