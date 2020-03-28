# KONGA Backend

[![wercker status](https://app.wercker.com/status/ba18dfbce20c686dc42e56ec32e400d9/s/master "wercker status")](https://app.wercker.com/project/byKey/ba18dfbce20c686dc42e56ec32e400d9)
[![Go Report Card](https://goreportcard.com/badge/gitlab.com/hyperd/konga-backend)](https://goreportcard.com/report/gitlab.com/hyperd/konga-backend)
[![GoDoc](https://godoc.org/gitlab.com/hyperd/konga-backend?status.svg)](https://pkg.go.dev/gitlab.com/hyperd/konga-backend?tab=doc)

Backend project for Konga, a full-fledged management interface for [Kong API Platform](https://konghq.com/).

## Build the Konga Backend

There are two ways here available to build the code; a targetted method and a [cross-platform builder script](./build.bash). Both allow to create portable executables, compatible with [Alpine Linux](https://www.alpinelinux.org/), compiled statically linking C bindings `-installsuffix 'static'`, and omitting the symbol and debug info `-ldflags "-s -w"`.

### Targetted build, based on your system/architecture

```bash
# change according to your system/architecture
CGO_ENABLED=0 GOARCH=[amd64|386] GOOS=[linux|darwin] go build -ldflags="-w -s" -a -installsuffix 'static' -o konga cmd/konga/main.go
```

### Multiplatform build with Docker

To trigger a multiplatform build, from the root diretory of this repository, run:

```bash
rm -rf $(pwd)/releases/*

docker run --rm -it -v "$PWD":/usr/src/app -w /usr/src/app golang:latest bash -c '
for GOOS in darwin linux; do
    for GOARCH in 386 amd64; do
      export GOOS GOARCH
      CGO_ENABLED=0 GO111MODULE=on go build -ldflags="-w -s -X main.minversion=`date -u +.%Y%m%d.%H%M%S`" \
      -a -installsuffix "static" -o releases/konga-$GOOS-$GOARCH cmd/konga/main.go
    done
done
'
```

#### Local testing with the inmemory strategy

For local testing run the code as follow:

```bash
./init-modules.bash && go run cmd/konga/main.go --database.type="inmemory"
```

Or execute the compiled binary:

```bash
./releases/konga-darwin-amd64 --database.type="inmemory"
```
