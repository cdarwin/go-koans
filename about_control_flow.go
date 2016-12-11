package go_koans

import "fmt"

func aboutControlFlow() {
	{
		a, b, c := 1, 2, 3
		assert(a == 1) // multiple assignment
		assert(b == 2) // can make
		assert(c == 3) // life easier
	}

	var str string

	{
		if 3.14 == 3 {
			str = "what is love?"
		} else {
			str = "baby dont hurt me"
		}
		assert(str == "baby dont hurt me") // no more

		if length := len(str); length == 17 {
			str = "to be"
		} else {
			str = "or not"
		}
		assert(str == "to be") // that is the question
	}

	{
		hola1, hola2 := "ho", "la"

		switch "hello" {
		case "hello":
			str = "hi"
		case "world":
			str = "planet"
		case fmt.Sprintf("%s%s", hola1, hola2):
			str = "senor"
		}
		assert(str == "hi") // cases can be of any type, even arbitrary expressions

		switch {
		case false:
			str = "first"
		case true:
			str = "second"
		}
		assert(str == "second") // in the absence of value, there is truth
	}

	{
		n := 0
		for i := 0; i < 5; i++ {
			n += i
		}
		assert(n == 10) // for can have the structure with which we are all familiar
	}

	{
		n := 1
		for {
			n *= 2
			if n > 20 {
				break
			}
		}
		assert(n == 32) // though omitting everything creates an infinite loop
	}
}
