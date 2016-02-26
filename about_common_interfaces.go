package go_koans

import "bytes"
import "io"

func aboutCommonInterfaces() {
	{
		in := new(bytes.Buffer)
		in.WriteString("hello world")

		out := new(bytes.Buffer)

        in.WriteTo(out)

		/*
		   Your code goes here.
		   Hint, use these resources:

		   $ godoc -http=:8080
		   $ open http://localhost:8080/pkg/io/
		   $ open http://localhost:8080/pkg/bytes/
		*/

		assert(out.String() == "hello world") // get data from the io.Reader to the io.Writer
	}

	{
		in := new(bytes.Buffer)
		in.WriteString("hello world")
        foo := io.LimitReader(in, 5)

		out := new(bytes.Buffer)
        out.ReadFrom(foo)

		assert(out.String() == "hello") // duplicate only a portion of the io.Reader
	}
}
