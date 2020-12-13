package go_koans

func aboutDefer() {
	var acc int

	increment := func(value int) {
		acc += value
	}

	decrement := func(value int) {
		acc -= value
	}

	panicRecover := func() {
		if r := recover(); r != nil {
			increment(1)
		}
	}

	func() {
		acc = 0
		defer increment(1)
	}()

	assert(acc == 1) // defer function will be execute after main function body

	func() {
		acc = 0
		defer increment(5)
		defer decrement(3)
	}()

	assert(acc == 2) // list of functions also allowed

	func() {
		defer panicRecover()
		acc = 0
		panic("Expected error")
	}()

	assert(acc == 1) // executed even in case of panic

}
