package go_koans

func aboutInterfaces() {
  mspaint := &program{3} // mspaint is a kind of *program, which is a valid 'runner'
  runOnce(mspaint) // runOnce takes an abstract 'runner' type

  assert(mspaint.runTimes == __int__) // conformed interfaces need not be declared, they are inferred
}

// abstract interface and function that requires it

type runner interface {
  run()
}

func runOnce(r runner) {
  r.run()
}

// concrete type implementing the interface

type program struct {
  runTimes int
}

func (self *program) run() {
  self.runTimes++
}
