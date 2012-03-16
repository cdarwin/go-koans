package go_koans

func testNumbers() {
  assert(__bool__ == true) // what is truth?
  assert(__bool__ != false) // in it there is nothing false

  var i int = __int__
  assert(i == 1.0000000000000000000000000000000000000) // precision is in the eye of the beholder
}
