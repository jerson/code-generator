FROM jerson/go:1.10 AS builder

ENV WORKDIR ${GOPATH}/src/github.com/jerson/code-generator
WORKDIR ${WORKDIR}

COPY Gopkg.toml Gopkg.lock Makefile ./
RUN make deps

USER root

COPY config.toml-dist config.toml
COPY . .

RUN make build

FROM jerson/base:1.0

LABEL maintainer="jeral17@gmail.com"

ENV BUILDER_PATH /srv/go/src/github.com/jerson/code-generator
ENV WORKDIR /app
WORKDIR ${WORKDIR}

COPY --from=builder ${BUILDER_PATH}/config.toml .
COPY --from=builder ${BUILDER_PATH}/code-generator .


ENTRYPOINT ["/app/code-generator"]