package go_koans

func aboutChannels() {
	ch := make(chan string, 2)

	assert(len(ch) == __int__) // channels are like buffers

	ch <- "foo" // i mean, "metaphors are like similes"

	assert(len(ch) == __int__) // they can be queried for queued items

	assert(<-ch == __string__) // items can be popped out of them

	assert(len(ch) == __int__) // and len() always reflects the "current" queue status

	// the 'go' keyword runs a function-call in a new "goroutine"
	// which executes "concurrently" with the calling "goroutine"
	go func() {
		// your code goes here
	}()

	assert(__delete_me__) // we'll need to make room for the queue, or suffer deadlocks

	ch <- "bar"   // this send will succeed
	ch <- "quux"  // there's enough room for this send too
	ch <- "extra" // but the buffer only has two slots
}
