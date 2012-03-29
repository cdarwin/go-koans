## Go Koans

Want to learn Go? Want to do it easily? Want to have fun with it? Want lots of money?

I can't help you out with that last one. Sorry.

### References

Bookmark the [spec](http://golang.org/ref/spec) and the
[packages listing](http://golang.org/pkg/). You can also
run the Go website locally with `godoc -http=:8080`.

### Usage

    $ go test

### Better usage

Annoyed at having to switch between your editor and terminal, just to type `go test` over and over?
As you should be! So download fswatch and let it run your tests for you any time you save a file.

    $ go get github.com/sdegutis/fswatch
    $ fswatch -f . -- go test

(If you get a 'permission denied' error on the 'go get' command and don't want to use sudo, do the
"Go Setup" steps below.)

fswatch is kind of like watchr, but it's a native executable, written in Go, and only runs on Macs
(you are using a Mac, right?) using the Apple-specific FSEvents framework. So it doesn't poll files
or anything inefficient like that. It's very healthy for your system. And it's not a gem and so
doesn't depend on any other programs, it's just a native compiled executable.

### Go Setup

After installing Go, these steps are necessary before any Go project will work. All
project directories need to be under `$GOPATH/src` for the `go` command to be useful.

    $ export GOPATH=~/my-go-projects
    $ mkdir -p $GOPATH/src

Since the `$GOPATH` is almost necessary for Go, it's recommended that this goes in your `~/.profile`:

    export GOPATH="$HOME/my-go-projects"
    export PATH="$GOPATH/bin:$PATH"

### Benefaxion

Anyway, diggin' it? If so, endorse me:

[![endorse](http://api.coderwall.com/sdegutis/endorse.png)](http://coderwall.com/sdegutis)

Thanks!
