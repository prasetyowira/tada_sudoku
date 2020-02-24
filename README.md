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
$ docker pull ariwira/tada_sudoku:latest
$ docker run --rm ariwira/tada_sudoku bash -c "tada_sudoku < /opt/tada_sudoku/solved.txt"
or
$ docker run --rm ariwira/tada_sudoku bash -c "tada_sudoku solve < /opt/tada_sudoku/input.txt"
```

### Testing

``make test``