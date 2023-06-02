# lucy

Lucy is a command line tool to send/stream files among devices on the same LAN over TCP

## Installation

download and install the Go programming language from this link: [download Go](https://go.dev/dl/)

then run:

```shell
go install github.com/notEpsilon/lucy@latest
```

and you are done!.

## Usage

on the receiving computer open a cmd/terminal and run:

```shell
lucy wait -o output_file.zip
```

> Note: the `output_file.zip` is the output file name you need to **PROVIDE CORRECT OUTPUT FILE EXTENSION**

then on the sender computer run:

```shell
lucy send -f my_file.zip
```

## Options

to see different options and default values run:

```shell
lucy help
```
