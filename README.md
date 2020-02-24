# tada_sudoku

A Sudoku Checker with Golang for TADA interview assignment

## Getting started

This project requires Go to be installed. On OS X with Homebrew you can just run `brew install go`.

Running it then should be as simple as:

```console
$ make
$ ./bin/tada_sudoku < solved.txt
```

To run sudoku solver
```console
$ ./bin/tada_sudoku solver < input.txt
```

To build inside docker
```console
$ make package
```

To run inside docker
```console
$ docker run ariwira/tada-sudoku-image:latest
```

### Testing

``make test``