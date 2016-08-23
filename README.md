# Go Koans

I'm not sure how this ever got popular under my namespace, but it seems to have
enough interest that I feel compelled to maintain it now. The original author,
[Steven Degutis](https://github.com/sdegutis), committed the initial suite of
tests, 4c5e766, on Mar 10, 2012. I don't recall now how I discovered the
initial codebase beyond searching for "go koans" on GitHub. I do recall that
I was enlightened considerably by [Ruby Koans](http://rubykoans.com/) and
something similarly enlightening on my journey to the Nirvana of Go could be
a blessing for anyone.

Since my discovery of [The Go Programming Language](https://golang.org/), the
language and development environments have changed significantly. I will do my
best to balance current best practices and a low barrier of entry for newcomers
(whom I assume to be the vast majority of those with interest in this
repository). I will try to keep up to date with the latest stable releases. I
hope I can rely on this wonderful community to help me with this.

## Native Usage

If you prefer to do things "natively" you may download, install, and configure
the suite of tools provided by the binary or source distribution of your
choice from the [Downloads page](https://golang.org/dl/).

1. Run `go test`.
1. Make the failing tests pass, by replacing these types of `__variables__` with real values.

## Docker Usage

I have found that using [Docker](https://www.docker.com/) helps me keep my
development environment clean and portable. Here is an example of how I might
set up an environment dedicated to go through these koans.

Install/Setup:

```shell
luser@lolcathost:~ $ docker-machine create -d virtualbox golang
luser@lolcathost:~ $ eval $(docker-machine env golang)
luser@lolcathost:~ $ docker pull library/golang:1.6.0-alpine
luser@lolcathost:~ $ docker run --rm -ti -v "$PWD":/usr/src/koans -w /usr/src/koans golang:1.6.0-alpine /bin/sh
```

Now with an interactive shell inside of a minimal container you may iterate
through the same steps to enlightenment described above.

## Helpful References

Bookmark the [spec](http://golang.org/ref/spec) and the [packages listing](http://golang.org/pkg/).
You can also run the Go website locally with `godoc -http=:8080`.

## Go support in Vim

If you have an interest in a more fancy vim setup, I urge you to consider
the post on the [Go Development Environment for Vim](https://blog.gopheracademy.com/vimgo-development-environment/)
on the Gopher Academy Blog.

## Go-Koans support in Emacs

[Jacek Wysocki](https://github.com/exu) has provided some nice Go Koans helper
scripts for Emacs users at [exu/go-koans.el](https://github.com/exu/go-koans.el)
