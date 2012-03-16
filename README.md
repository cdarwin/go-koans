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

### Usage

    $ go test

### More fun usage

    $ brew install https://raw.github.com/sdegutis/fswatch/master/brew_formula/fswatch.rb
    $ fswatch -f . go test
