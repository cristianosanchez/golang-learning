# Go

Install Go, downloading from site or Homebrew.

  ```$ brew install go```

Default `$GOPATH` is `$HOME/go`, you may change that location poitinng to somewhere else. Check go environment with `go env` for more.

Note that your project might though be using go modules. Check `go.mod` file from your project root folder.

All projects will be located at `$GOPATH/src/github.com/cristianosanchez/<project>`. Commands will be installed at `$GOPATH/bin`. The `$GOPATH/pkg` will hold builded packages.

You don't need to be at directory to run commands, so, from anywhere, this would work:

  ```$ go get github.com/digitalocean/doctl```

## Go CLI

Type `go help` for a list of commands, mainly:

. build
. clean
. doc
. env
. bug
. fix
. fmt
. generate
. get
. install
. list
. run
. test
. tool
. version
. vet
