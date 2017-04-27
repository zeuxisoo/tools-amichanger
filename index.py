#!/usr/bin/env python3
# -*- coding: utf-8 -*-

import os
import sys
import random
import subprocess
import hashlib

import click

__VERSION__ = "0.1.0"

program_root = os.path.dirname(os.path.abspath(__file__))
tools_root   = os.path.join(program_root, "tools")
results_root = os.path.join(program_root, "results")

amiitool = os.path.join(tools_root, "amiitool")

def md5_file(file_path):
    return hashlib.md5(open(file_path.replace('\\', ''), 'rb').read()).hexdigest()

def generate_serial(*args):
    return "{:02X}{:02X}{:02X}{:02X}{:02X}{:02X}{:02X}{:02X}{:02X}".format(*args)

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

    # Generate serial
    uid0 = random.randint(0, 255)
    uid1 = random.randint(0, 255)
    uid2 = random.randint(0, 255)
    uid3 = random.randint(0, 255)
    uid4 = random.randint(0, 255)
    uid5 = random.randint(0, 255)
    uid6 = random.randint(0, 255)

    bcc0 = uid0 ^ uid1 ^ uid2
    bcc1 = uid3 ^ uid4 ^ uid5 ^ uid6

    serial = generate_serial(uid0, uid1, uid2, bcc0, uid3, uid4, uid5, uid6, bcc1)

    click.echo("Generated serial : {}".format(serial))
    click.echo("Generated values : uid0={uid0} uid1={uid1} uid2={uid2} bcc0={bcc0} uid3={uid3} uid4={uid4} uid5={uid5} uid6={uid6} bcc1={bcc1}".format(
        uid0=uid0, uid1=uid1, uid2=uid2, bcc0=bcc0, uid3=uid3, uid4=uid4, uid5=uid5, uid6=uid6, bcc1=bcc1
    ))
    click.echo("-----------")

    # Decrypt amiibo
    decrypt_filename = "{0}_decrypt.bin".format(os.path.splitext(os.path.basename(amiibo))[0])
    decrypt_filepath = os.path.join(results_root, decrypt_filename)
    decrypt_command  = "{amiitool} -d -k {key} -i {amiibo} -o {decrypt_filepath}".format(
        amiitool=amiitool, key=key, amiibo=amiibo, decrypt_filepath=decrypt_filepath
    )

    click.echo("Start decrypting ...")
    click.echo("=> {}".format(decrypt_command))

    shell = subprocess.Popen(decrypt_command, shell=True)
    stdout, stderr = shell.communicate()

    if stderr is None:
        click.echo("=> MD5: {}".format(md5_file(decrypt_filepath)))
        click.echo("=> OK")
    else:
        click.echo(stderr)
        click.echo("=> Failed")

        raise SystemExit(0)

    # Change amiibo serial
    click.echo("Start changing ...")

    try:
        with open(decrypt_filepath, 'r+b') as fh:
            fh.seek(0, 0)
            fh.write(bytes([bcc1]))
            fh.seek(0x1D4, 0)
            fh.write(bytes([bcc0, uid0, uid1, uid2, uid3, uid4, uid5, uid6]))
            fh.close()

        click.echo("=> MD5: {}".format(md5_file(decrypt_filepath)))
        click.echo("=> OK")
    except Exception as e:
        click.echo(e)
        click.echo("=> Failed")

        raise SystemExit(0)

    # Encrypt changed amiibo decrypt bin
    encrypt_filename = "{0}_{1}.bin".format(os.path.splitext(os.path.basename(amiibo))[0], serial)
    encrypt_filepath = os.path.join(results_root, encrypt_filename)
    encrypt_command  = "{amiitool} -e -k {key} -i {decrypt_filepath} -o {encrypt_filepath}".format(
        amiitool=amiitool, key=key, decrypt_filepath=decrypt_filepath, encrypt_filepath=encrypt_filepath
    )

    click.echo("Start encrypting ...")
    click.echo("=> {}".format(encrypt_command))

    shell = subprocess.Popen(encrypt_command, shell=True)
    stdout, stderr = shell.communicate()

    if stderr is None:
        click.echo("=> MD5: {}".format(md5_file(encrypt_filepath)))
        click.echo("=> OK")
    else:
        click.echo(stderr)
        click.echo("=> Failed")

        raise SystemExit(0)

if __name__ == '__main__':
    if sys.version_info < (3,0,0):
        print("\nDon\'t not support python 2, Please run on python 3\n")
    else:
        run()
