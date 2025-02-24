#!/bin/bash
CURDIR=$(cd $(dirname $0); pwd)
BinaryName=fronted
echo "$CURDIR/bin/${BinaryName}"
exec $CURDIR/bin/${BinaryName}