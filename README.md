# samspackagefun

This is a demo project in go that generates multiple binaries that use shared functions. 


## Install and run

    $ go get github.com/samjbobb/samspackagefun/...
    
Which will retrieve the library, compile, and install the `samfoo` and `sambar` binaries in your `$GOBIN` 
which is usually `$GOPATH/bin`.

Then run:

    $ samfoo
    $ sambar
    
## Development

Re-compile those binaries into `$GOBIN`:
     
     $ go install github.com/samjbobb/samspackagefun/...
     
     
Build both into another dir??? ...


Build `samfoo` into current dir:

    $ go build ./cmd/samfoo


## Reading

[Structuring Applications in Go](https://medium.com/@benbjohnson/structuring-applications-in-go-3b04be4ff091#.t8p2uhaol)

https://groups.google.com/forum/#!topic/golang-nuts/mAtDVVaUk2c

https://golang.org/cmd/go/#hdr-Description_of_package_lists

> An import path that is a rooted path or that begins with a . or .. element is interpreted as a file system path and denotes the package in that directory.

> Otherwise, the import path P denotes the package found in the directory DIR/src/P for some DIR listed in the GOPATH environment variable (see 'go help gopath').

> An import path is a pattern if it includes one or more "..." wildcards, each of which can match any string, including the empty string and strings containing slashes. Such a pattern expands to all package directories found in the GOPATH trees with names matching the patterns. As a special case, x/... matches x as well as x's subdirectories. For example, net/... expands to net and packages in its subdirectories.



## Other (better) examples

https://github.com/boltdb/bolt

https://github.com/camlistore/camlistore

More pro.
