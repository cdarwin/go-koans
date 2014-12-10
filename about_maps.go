package go_koans

func aboutMaps() {
	ages := map[string]int{
		"bob": 10,
		"joe": 20,
		"dan": 30,
	}

	age := ages["bob"]
	assert(age == __int__) // map syntax is warmly familiar

	age, ok := ages["bob"]
	assert(ok == __bool__) // with a handy multiple-assignment variation

	age, ok = ages["steven"]
	assert(age == __int__)    // the zero value is used when absent
	assert(ok == __boolean__) // though there are better ways to check for presence

	assert(len(ages) == __int__) // length is based on keys

	ages["bob"] = 99
	assert(ages["bob"] == __int__) // values can be changed for keys

	ages["steven"] = 77
	assert(ages[__string__] == 77) // new ones can be added

	delete(ages, "steven")
	age, ok = ages["steven"]
	assert(ok == __boolean__) // key/value pairs can be removed
}
