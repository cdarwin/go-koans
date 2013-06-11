## Go Koans

Want to learn Go? Want to do it easily? Want to have fun with it? Want lots of money?

I can't help you out with that last one. Sorry.

### Usage

  1. [Install Go](http://code.google.com/p/go/downloads/list).
  2. Clone this repo.
  3. Run `go test`.
  4. Make the failing tests pass, by replacing these types of `__variables__` with real values.

#### Cooler usage

If you want to have `go test` be run in your terminal any time you save a file and
are using a Mac, take a look at [fswatch](http://github.com/sdegutis/fswatch).

### Helpful References

Bookmark the [spec](http://golang.org/ref/spec) and the
[packages listing](http://golang.org/pkg/). You can also
run the Go website locally with `godoc -http=:8080`.

### Go support in Vim

Add this to your `~/.vimrc` file:

    set rtp+=/usr/local/go/misc/vim
    filetype plugin indent on
    syntax on

### Go support in Emacs

To install Go mode for emacs check this stack response http://stackoverflow.com/questions/1715464/emacs-lisp-mode-for-go

Go-koans have Emacs function to speed up koans resolvig. Simply put `go-koans.el` file somewhere in your `load-path`
and then

    (require 'go-koans)

into your `init.el` file (helper integrates with great Expand region extension from magnars: https://github.com/magnars/expand-region.el)

next `cd` into your koans directory and press `C-c C-r` - Emacs will open file and goto line where error occurs


### Benefaxion

Anyway, diggin' it? If so, endorse me:

[![endorse](http://api.coderwall.com/sdegutis/endorse.png)](http://coderwall.com/sdegutis)

Thanks!
