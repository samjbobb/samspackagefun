# samspackagefun

This is a demo project in go that generates multiple binaries that use shared functions. 


## Install and run

    $ go get github.com/samjbobb/samspackagefun/...
    
Which will retreive the library and install the `samfoo` and `sambar` executables in your `$GOBIN` 
which is usually `$GOPATH/bin`.

Then run:

    $ samfoo
    $ sambar
    
## Development

Re compile those executables into the same dir:
     
     $ go install github.com/samjbobb/samspackagefun/...
     
Build into another dir???


## Reading

[Structuring Applications in Go](https://medium.com/@benbjohnson/structuring-applications-in-go-3b04be4ff091#.t8p2uhaol)

https://groups.google.com/forum/#!topic/golang-nuts/mAtDVVaUk2c

https://golang.org/cmd/go/#hdr-Description_of_package_lists

## Other (better) examples

https://github.com/boltdb/bolt

https://github.com/camlistore/camlistore

More pro.
