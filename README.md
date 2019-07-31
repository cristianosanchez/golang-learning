# GoLang

Install Go, downloading from site or Homebrew.

  brew install go

Default $GOPATH is $HOME/go, you may change that location.

All projects than will to be located at (example) `$GOPATH/src/github.com/cristianosanchez/<repo>` and all dependencies will be installed at `$GOPATH/bin`. That's the default behaviour before `go modules`.

You don't need to be at directory to run commands, so, from anywhere, this would work:

  go get github.com/digitalocean/doctl

Try go modules and be free of GOPATH dictatorship.

# Go CLI

Type ``` go help``` for a list of commands, mainly: ```build, clean, doc, env, bug, fix, fmt, generate, get, install, list, run, test, tool, version, vet```

