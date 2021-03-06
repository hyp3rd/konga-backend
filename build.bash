#!/bin/bash
set -Eeuo pipefail

traperr() {
  echo "ERROR: ${BASH_SOURCE[1]} at about line ${BASH_LINENO[0]} ${ERR}"
}

set -o errtrace
trap traperr ERR

report () {
	cat >&2 <<-'EOF'

The Docker images are now up to date;
remember to change the sha256 checksum in the k8s deployment file.

	EOF
}

validate_env () {
  if [[ -z ${PROJECT_ID+x} ]] || [[ -z ${REGION+x} ]]; then
    echo "To run this deployment you need to export PROJECT_ID and REGION as follows:
    export REGION=<region e.g. europe-west1>
    export PROJECT_ID=<project name e.g. hyperd-konga-backend>";
    exit 1
  fi
}

# build it's a simple golang cross-compiler that generates alpine linux compatible binaries
build () {

  rm -rf "$(pwd)"/releases/*

  docker run --rm -it -v "$PWD":/usr/src/app -w /usr/src/app golang:latest bash -c '
  for GOOS in darwin linux; do
      for GOARCH in 386 amd64; do
        export GOOS GOARCH
        CGO_ENABLED=0 GO111MODULE=on go build -ldflags="-w -s -X main.minversion=`date -u +.%Y%m%d.%H%M%S`" \
        -a -installsuffix "static" -o releases/konga-$GOOS-$GOARCH cmd/konga/main.go
      done
  done
  '
}

push_to_scm() {
  git status
  git add . ; git commit -m "Automated build" || : ; git push || :
}

update_docker_images () {
  validate_env

  cd "$(pwd)"/releases

  # build the konga-backend image with our modifications (see Dockerfile) and tag for private GCR
  docker build --no-cache --file ../.containers/Dockerfile -t gcr.io/$PROJECT_ID/konga-backend:latest .

  # configure pushing to private GCR, and push our image
  gcloud auth configure-docker -q
  docker push gcr.io/"$PROJECT_ID"/konga-backend:latest
  docker push gcr.io/"$PROJECT_ID"/konga-backend:mock

  report
}

build && push_to_scm && update_docker_images
