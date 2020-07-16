# go-algorithm

## Description

This repository is for learning purpose, containing algorithms.

## Dependency

### GO Language

The GO programming language version `go1.14.4` need to be installed
in the system. Go to [this link](https://golang.org/doc/install) and follow
the instruction to install. To check GO version on the environment, run the following command:

```bash
go version
```

example output:

```bash
go version go1.14.4 linux/amd64
```

### Go Dep

The Dep, for dependency management tool for Go, version `devel` should be installed.
Go to [this link](https://github.com/golang/dep) and follow the instruction to install based on the system.
To check the installation, we can check its version by running the following command on the terminal:

```bash
dep version
```

Example of the output:

```bash
dep:
 version     : devel
 build date  : 
 git hash    : 
 go version  : go1.14.4
 go compiler : gc
 platform    : linux/amd64
 features    : ImportDuringSolve=false
```

## How to Test

In this project root directory, run the following command to test:

```bash
make test
```

To show a more complete coverage and uncovered lines:

```bash
make test_coverage
```

You can check into `test_coverage.html` file in root project directory. This command also will open interactive coverage tool in your browser if you have one.

## How to Build

In this project root directory, run the following command:

```bash
make build
```

There will be a new directory named `out` with an executable file `go-algorithm` as the result of the built project.

## How to Run

In order to run this project, after building this project, execute
the following command in this project root directory:

```bash
make run
```
