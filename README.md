# Tools AmiChanger (Python)

A tools for change the serial number in dump ami file

## Installation

Create the virtual environment and install the required packages

    make venv

Generate the important tools

    make tools

Clean all generated and created files

    make clean

## Usage

Switch to virtual environment

    source venv/bin/activate

Generate single changed file

    python index.py create --key=/path/to/your/retail/key.bin --amiibo=/path/to/your/dump/ami.bin

Generate multiple changed file (sample: 3 files)

    python index.py multi --key==/path/to/your/retail/key.bin --amiibo=/path/to/your/dump/ami.bin --count=3

Debug mode

    python index.py --debug [create|multi] [options]

