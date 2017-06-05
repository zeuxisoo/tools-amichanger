# Tools AmiChanger

A tools for change the serial number in dump ami file

## Prepare

Install the vendors

    glide install

Download the tools packages

    make amiitool

## Build

Build the program

    make build

Show help message from generated file

    ./main

## Develop

Build the program, and try to generate single file

    go run *.go create --key=/path/to/retail/key.bin --amiibo=/path/to/ami/dump/file.bin

Build the program, and try to generate multiple files base on count arguments

    go run *.go multi --key=/path/to/retail/key.bin --amiibo=/path/to/ami/dump/file.bin --count=10

Clean the generated bin files

    make clean-results

## Other versions

| Language     | Link                        |
| ------------ | --------------------------- |
| Python       | [here][python version link] |

[python version link]: https://github.com/zeuxisoo/tools-amichanger/tree/python
