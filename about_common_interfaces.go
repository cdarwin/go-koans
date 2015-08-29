package go_koans

import "bytes"

func aboutCommonInterfaces() {
	{
		in := new(bytes.Buffer)
		in.WriteString("hello world")

		out := new(bytes.Buffer)

		/*
		   Your code goes here.
		   Hint, use these resources:

		   $ godoc -http=:8080
		   $ open http://localhost:8080/pkg/io/
		   $ open http://localhost:8080/pkg/bytes/
		*/

		// get data from the io.Reader to the io.Writer
		in.WriteTo(out)

		assert(out.String() == "hello world")
	}

	{
		in := new(bytes.Buffer)
		in.WriteString("hello world")

		out := new(bytes.Buffer)
		in.Truncate(5)
		in.WriteTo(out)

		assert(out.String() == "hello") // duplicate only a portion of the io.Reader
	}
}
