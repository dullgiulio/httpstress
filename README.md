## httpstress
httpstress is a Go library for HTTP stress testing.
It launches one goroutine per concurrent connection.

A CLI is provided in [httpstress-go repo](https://github.com/chillum/httpstress-go.git).
If you need a CLI utility, install [httpstress-go](https://github.com/chillum/httpstress-go.git),
not httpstress.

### Installation
* Install [Git](http://git-scm.com/download)
* Install [Go runtime](http://golang.org/doc/install).
  Go 1.3 or higher on amd64 is recommended because of performance issues
* Set [`GOPATH`](http://golang.org/doc/code.html#GOPATH)
* `go get github.com/chillum/httpstress`

### Docs
* [godoc.org](https://godoc.org/github.com/chillum/httpstress)
* `godoc github.com/chillum/httpstress`
