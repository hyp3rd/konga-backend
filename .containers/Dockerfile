#+++++++++++++++++++++++++++++++++++++++++++++++++++++++++
# Dockerfile for eu.gcr.io/hyperd-containers/konga-backend
#+++++++++++++++++++++++++++++++++++++++++++++++++++++++++

FROM alpine:3.11.5

ENV TERM="xterm" \
    LANG="C.UTF-8" \
    LC_ALL="C.UTF-8" \
    LOG_STDOUT="" \
    LOG_STDERR=""

RUN set -x \
  && apk update \
  && rm -rf /var/cache/apk/* \
  export RANDOM_PASSWORD=`tr -dc A-Za-z0-9 < /dev/urandom | head -c71` \
  && echo "root:$RANDOM_PASSWORD" | chpasswd \
  && unset RANDOM_PASSWORD \
  && passwd -l root

RUN set -ex; \
	apkArch="$(apk --print-arch)"; \
	case "$apkArch" in \
	armhf) arch='armv6' ;; \
	aarch64) arch='arm64' ;; \
	x86_64) arch='amd64' ;; \
	*) echo >&2 "error: unsupported architecture: $apkArch"; exit 1 ;; \
	esac; \
	wget --quiet -O /usr/local/bin/konga "https://gitlab.com/hyperd/konga-backend/raw/master/releases/konga-linux-$arch"; \
	chmod +x /usr/local/bin/konga

RUN adduser -D -g '' appuser

USER appuser

EXPOSE 3000 8443

ENTRYPOINT ["konga"]

# Metadata
LABEL org.opencontainers.image.vendor="Hyperd" \
	org.opencontainers.image.url="https://hyperd.sh" \
	org.opencontainers.image.title="Konga Backend" \
	org.opencontainers.image.description="Backend project for Konga, a full-fledged management interface for Kong API Platform" \
	org.opencontainers.image.version="v0.0.1" \
	org.opencontainers.image.documentation="https://gitlab.com/hyperd/konga-backend/blob/master/README.md"
