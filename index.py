#!/usr/bin/env python3
# -*- coding: utf-8 -*-

import sys
import click

from changer import commands
from changer import utils

@click.group()
@click.option('--debug/--no-debug', default=False)
@click.pass_context
def cli(ctx, debug):
    ctx.obj['DEBUG'] = debug

@cli.command()
@click.pass_context
@click.option("--key", required=True, help="The path of retail bin file")
@click.option("--amiibo", required=True, help="The path of amiibo dump file")
def create(ctx, key, amiibo):
    """Create new amiibo file with random serial"""

    commands.create(ctx, key, amiibo)


if __name__ == '__main__':
    if sys.version_info < (3,0,0):
        utils.raise_message("Don\'t not support python 2, Please run on python 3")
    else:
        cli(obj={})
