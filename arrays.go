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
}
