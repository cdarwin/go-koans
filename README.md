## Go Koans

Want to learn Go? Want to do it easily? Want to have fun with it? Want lots of money?

I can't help you out with that last one. Sorry.

### Go Setup

After installing Go, these steps are necessary before any Go project will work. All
project directories need to be under `$GOPATH/src` for the `go` command to be useful.

    $ export GOPATH=~/my-go-projects
    $ mkdir -p $GOPATH/src

Since the `$GOPATH` is useful any time in Go, this is recommended:

    $ echo 'export GOPATH=~/my-go-projects' >> ~/.profile

### Install the Koans

    $ cd $GOPATH/src
    $ git clone https://sdegutis@github.com/sdegutis/go-koans.git
    $ cd go-koans

### References

Bookmark the [spec](http://weekly.golang.org/ref/spec) and the
[packages listing](http://weekly.golang.org/pkg/). You can also
run the Go website locally with `godoc -http=:8080`.

### Usage

    $ go test

### More fun usage

    $ brew install --HEAD https://raw.github.com/sdegutis/fswatch/master/brew_formula/fswatch.rb
    $ fswatch -f . go test

### Benefaxion

Anyway, diggin' it? If so, endorse me:

[![endorse](http://api.coderwall.com/sdegutis/endorse.png)](http://coderwall.com/sdegutis)

Thanks!
