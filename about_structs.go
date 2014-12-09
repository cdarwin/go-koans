package go_koans

func aboutStructs() {
	var bob struct {
		name string
		age  int
	}
	bob.name = "bob"
	bob.age = 30

	assert(bob.name == __string__) // structs are collections of named variables
	assert(bob.age == __int__)     // each field has both setter and getter behavior

	type person struct {
		name string
		age  int
	}

	var john person
	john.name = "bob"
	john.age = __int__

	assert(bob == john) // assuredly, bob is certainly not john.. yet
}
