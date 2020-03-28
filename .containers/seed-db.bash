#!/bin/bash

set -Eeuo pipefail

traperr() {
  echo "ERROR: ${BASH_SOURCE[1]} at about line ${BASH_LINENO[0]} ${ERR}"
}

set -o errtrace
trap traperr ERR

seed() {
	# TODO: implement a more elegant way to wait for the cluster
	# to be up and running.
    echo "Wait for servers to be up"
    sleep 10

    HOSTPARAMS="--host roach1 --insecure"
    SQL="/cockroach/cockroach.sh sql $HOSTPARAMS"

    $SQL -e "CREATE DATABASE konga; CREATE USER IF NOT EXISTS hyperd; GRANT ALL ON DATABASE konga TO hyperd;"
}

seed
