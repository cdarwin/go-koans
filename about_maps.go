package go_koans

func aboutMaps() {
	ages := map[string]int{
		"bob": 10,
		"joe": 20,
		"dan": 30,
	}

	age := ages["bob"]
	assert(age == 10) // map syntax is warmly familiar

	age, ok := ages["bob"]
	assert(ok == true) // with a handy multiple-assignment variation

	age, ok = ages["steven"]
	assert(age == 0)    // the zero value is used when absent
	assert(ok == false) // though there are better ways to check for presence

	assert(len(ages) == 3) // length is based on keys

	ages["bob"] = 99
	assert(ages["bob"] == 99) // values can be changed for keys

	ages["steven"] = 77
	assert(ages["steven"] == 77) // new ones can be added

	delete(ages, "steven")
	age, ok = ages["steven"]
	assert(ok == false) // key/value pairs can be removed
}
