package go_koans

func divideFourBy(i int) int {
  return 4 / i
}

func aboutPanics() {
  n := divideFourBy(0)
  assert(n == 2) // panics are exceptional errors at runtime
}
