import click

from .create import create

def multi(ctx, key, amiibo, count):
    for i in range(count):
        click.echo("Creating number ---> {}".format(i))

        create(ctx, key, amiibo)

        click.echo("\n")
