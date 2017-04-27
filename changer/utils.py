import hashlib
import click

def md5_file(file_path):
    return hashlib.md5(open(file_path.replace('\\', ''), 'rb').read()).hexdigest()

def generate_serial(*args):
    return "{:02X}{:02X}{:02X}{:02X}{:02X}{:02X}{:02X}{:02X}{:02X}".format(*args)

def raise_message(message):
    click.echo("\n{0}\n".format(message))
