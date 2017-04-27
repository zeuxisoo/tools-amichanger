#!/usr/bin/env bash

ERROR_STATUS=0

function install {
    cd tools

    git clone https://github.com/socram8888/amiitool.git amiitool-src
    cd amiitool-src
    git submodule init
    git submodule update --recursive
    make
    cp amiitool ../amiitool
    cd ..
    rm -rf amiitool-src
}

function uninstall {
    cd tools
    rm -rf ./amiitool
}

COMMAND=${@:$OPTIND:1}

case $COMMAND in

    install)
        install
    ;;

    uninstall)
        uninstall
    ;;

    *)
        if [[ COMMAND != "" ]]; then
            print "Error: unknown command > $COMMAND\n\n"
            ERROR_STATUS=1
        fi

        usage
    ;;
esac

exit $ERROR_STATUS
