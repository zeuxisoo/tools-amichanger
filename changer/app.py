import os
import hashlib
import click

__VERSION__ = "0.1.0"

program_root = os.path.dirname(os.path.dirname(os.path.abspath(__file__)))
tools_root   = os.path.join(program_root, "tools")
results_root = os.path.join(program_root, "results")

amiitool = os.path.join(tools_root, "amiitool")
