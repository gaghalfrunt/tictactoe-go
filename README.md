# Tic Tac Toe in Go

## Introduction

A Tic Tac toe implementation in Go using a console user interface.

Uses [`gpm`](https://github.com/pote/gpm) for dependency management.
For easier `$GOPATH` and [Workspace](https://golang.org/doc/code.html#Workspaces) handling [`gvp`](https://github.com/pote/gvp) is recommended.


## Setup

###Requirements
- Go - [Download](https://golang.org/doc/install) (tested and developed with Go 1.4)
- [`gpm` &ndash; Go Package Manager](https://github.com/pote/gpm)
- [`gvp` &ndash; Go Versioning Packager](https://github.com/pote/gvp)

### Initial Setup

To set up the basic environment, please follow this steps.

1. The first step is to install the Go language and environment.
   Please do so by referring to the [official Go website](https://golang.org/doc/install)
   and follow their instructions for downloading and installing.

    Alternatively, Go can be installed with via Homebrew, too (e.g. `brew install go`).

2. After Go is installed, a tool called `gpm` (Go Package Manager) is needed.
   [The documentation of it](https://github.com/pote/gpm) mentions Homebrew as the
   preferred installation method.  

   This is done via your terminal by executing:

   `brew install gpm`

3. Hand in hand with `gpm` goes a tool called `gvp` (Go Versioning Packager).
   Its purpose is to ease the handling of one or more [Go Workspaces](https://golang.org/doc/code.html#Workspaces).  
   [The documentation of that tool](https://github.com/pote/gvp) mentions Homebrew as the
   preferred installation method, too.  

   Please do so by executing the following installation command in your terminal:

   `brew install gvp`


### Cloning or Downloading this Repository

If you're still in the terminal, the easiest way to download this repository is by cloning it to a local directory on your machine.

This can be done by executing
```
git clone https://github.com/gaghalfrunt/tictactoe-go.git
```

or by [downloading the zip archive](https://github.com/gaghalfrunt/tictactoe-go/archive/master.zip) and extracting it to a destination directory of your choice.


## Play the Game

If you're not in the terminal yet, now is the time to do so. Navigate to the directory where this repository has been downloaded or cloned.

### Setting up the Environment

Before running the game, make sure to download and install all dependencies, by executing:
```
gpm install
```

Then set the current workspace for Go with
```
source gvp in
```

Now everything is set for the environment and dependencies.

### Run the Game

For starting the game a [`main` package](https://golang.org/doc/code.html#Command) is available at `src/tictactoe.go` that can be directly executed with

```
go run src/tictactoe.go
```

Alternatively this file can be compiled into an executable and installed into the current Workspace via

```
go install src/tictactoe.go
```

Now the file `tictactoe` is available in the `$PATH` and can be executed directly.


## Run the Tests

For testing, [Ginkgo](http://onsi.github.io/ginkgo/) is used.

It comes with a custom executable to run the tests. To enable it please execute the two following commands in order to install Ginkgo in your environment:
```
go get github.com/onsi/ginkgo/ginkgo
go get github.com/onsi/gomega
```

After the dependencies have been installed, the `ginkgo` command is available on your system.

Executing all available suites/tests is done by
```
ginkgo -r
```
which will look recursively (`-r`) in all folders for test suites to execute.

Ginkgo integrates with the default Go testing tool `go test`. This can also be used to execute the tests:

```
go test -v tictactoe_test/...
```

Mind the three dots. It's a wildcard available for all `go` tools which basically says [&ldquo;this package and all its sub-packages&rdquo;](https://golang.org/cmd/go/#hdr-Relative_import_paths). The verbose flag (`-v`) is there so that you can actually see some output. Otherwise `go test` behaves relatively reticent and will only report `ok` as the result.


## Links
- [Tic Tac Toe in Go (this repository)](https://github.com/gaghalfrunt/tictactoe-go)
- [Go (language)](https://golang.org/doc/install)
- [`gpm` &ndash; Go Package Manager](https://github.com/pote/gpm)
- [`gvp` &ndash; Go Versioning Packager](https://github.com/pote/gvp)
- [Ginkgo - BDD Testing Framework for Go](http://onsi.github.io/ginkgo/)
- [Homebrew](http://brew.sh/)
