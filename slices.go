package go_koans

func testSlices() {
  fruits := []string{"apple", "orange", "mango"}

  assert(fruits[0] == __string__) // slices seem like arrays
  assert(len(fruits) == __int__) // in nearly all respects

  tasty_fruits := fruits[1:3] // we can even slice slices
  assert(tasty_fruits[0] == __string__) // slices of slices also share some data
  assert(tasty_fruits[1] == __string__) // but once again slightly askewed

  tasty_fruits[0] = "lemon"

  assert(fruits[0] == __string__) // but once again
  assert(fruits[1] == __string__) // some things have changed
  assert(fruits[2] == __string__) // since our data is not our own
}
