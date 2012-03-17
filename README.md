## Go Koans

Want to learn Go? Want to do it easily? Want to have fun with it? Want lots of money?

I can't help you out with that last one. Sorry.

### Go Setup

This part is specific to go, not really the koans. All project directories need
to be under `$GOPATH/src` for `go` to see them and do all it's awesome Go-stuff.

    $ export GOPATH=~/my-go-projects
    $ mkdir -p $GOPATH/src
    $ cd $GOPATH/src

Since the `$GOPATH` is useful any time in Go, this is recommended:

    $ echo 'export GOPATH=~/my-go-projects' >> ~/.profile

### Install the Koans

    $ git clone https://sdegutis@github.com/sdegutis/go-koans.git
    $ cd go-koans

### Getting Help

Two super useful references are the [spec](http://weekly.golang.org/ref/spec) and the [packages](http://weekly.golang.org/pkg/) listing.

You can also run the Go website locally with `godoc -http=:8080` so you don't need the internet. Then open [http://localhost:8080/ref/spec](http://localhost:8080/ref/spec) and [http://localhost:8080/pkg/](http://localhost:8080/pkg/).

### Usage

    $ go test

### More fun usage

    $ brew install --HEAD https://raw.github.com/sdegutis/fswatch/master/brew_formula/fswatch.rb
    $ fswatch -f . go test
