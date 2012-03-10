## Go Koans

Want to learn Go? Want to do it easily? Want to have fun with it? Want lots of money?

I can't help you out with that last one. Sorry.

### Install

    $ export GOPATH=~/projects/go # and add this to your ~/.profile
    $ mkdir -p ~/projects/go/src
    $ cd ~/projects/go/src

    $ git clone https://sdegutis@github.com/sdegutis/go-koans.git
    $ cd go-koans

### Usage

    $ go test

### More fun usage

    $ brew install https://raw.github.com/sdegutis/fswatch/master/brew_formula/fswatch.rb
    $ fswatch . go test
