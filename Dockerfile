# Build Stage
FROM lacion/alpine-golang-buildimage:1.13 AS build-stage

LABEL app="build-go-lak"
LABEL REPO="https://github.com/amyu98/go-lak"

ENV PROJPATH=/go/src/github.com/amyu98/go-lak

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:$GOROOT/bin:$GOPATH/bin

ADD . /go/src/github.com/amyu98/go-lak
WORKDIR /go/src/github.com/amyu98/go-lak

RUN make build-alpine

# Final Stage
FROM lacion/alpine-base-image:latest

ARG GIT_COMMIT
ARG VERSION
LABEL REPO="https://github.com/amyu98/go-lak"
LABEL GIT_COMMIT=$GIT_COMMIT
LABEL VERSION=$VERSION

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:/opt/go-lak/bin

WORKDIR /opt/go-lak/bin

COPY --from=build-stage /go/src/github.com/amyu98/go-lak/bin/go-lak /opt/go-lak/bin/
RUN chmod +x /opt/go-lak/bin/go-lak

# Create appuser
RUN adduser -D -g '' go-lak
USER go-lak

ENTRYPOINT ["/usr/bin/dumb-init", "--"]

CMD ["/opt/go-lak/bin/go-lak"]
