# Tools AmiChanger (Golang) (WIP)

A tools for change the serial number in dump ami file

## Prepare

Install the vendors

    glide install

Download the tools packages

    make ami

## Build

Build the program

    make build

Show help message from generated file

    ./main

## Develop

Build the program, and try to generate single file

    go run *.go create --key=/path/to/retail/key.bin --amiibo=/path/to/ami/dump/file.bin

## Other versions

| Language     | Link                        |
| ------------ | --------------------------- |
| Python       | [here][python version link] |

[python version link]: https://github.com/zeuxisoo/tools-amichanger/tree/python
