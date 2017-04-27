#!/usr/bin/env python3
# -*- coding: utf-8 -*-

import click

__VERSION__ = "0.1.0"

@click.command()
@click.option("--key", help="The path of retail bin file")
@click.option("--amiibo", help="The path of amiibo dump file")
def run(key, amiibo):
    click.echo("Changer")
    click.echo("-----------")
    click.echo("Current version  : {}".format(__VERSION__))
    click.echo("Your key path    : {}".format(key))
    click.echo("Your amiibo path : {}".format(amiibo))
    click.echo("-----------")

    click.echo("Decrypting ...")

    click.echo("Changing ...")

    click.echo("Encrypting ...")

if __name__ == '__main__':
    run()
