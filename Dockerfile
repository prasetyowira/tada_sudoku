# Build Stage
FROM ariwira/tada-sudoku-buildimage:1.13 AS build-stage

LABEL app="build-tada_sudoku"
LABEL REPO="https://github.com/prasetyowira/tada_sudoku"

ENV PROJPATH=/go/src/github.com/prasetyowira/tada_sudoku

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:$GOROOT/bin:$GOPATH/bin

ADD . /go/src/github.com/prasetyowira/tada_sudoku
WORKDIR /go/src/github.com/prasetyowira/tada_sudoku

RUN make build-alpine

# Final Stage
FROM alpine:latest

RUN apk add --no-cache --update \
    dumb-init \
    bash \
    ca-certificates \
    openssl && \
    rm -rf /var/cache/apk/*

ARG GIT_COMMIT
ARG VERSION
LABEL REPO="https://github.com/prasetyowira/tada_sudoku"
LABEL GIT_COMMIT=$GIT_COMMIT
LABEL VERSION=$VERSION

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:/opt/tada_sudoku/bin

WORKDIR /opt/tada_sudoku/bin

COPY --from=build-stage /go/src/github.com/prasetyowira/tada_sudoku/bin/tada_sudoku /opt/tada_sudoku/bin/
COPY ./input.txt tada_sudoku/
COPY ./solved.txt tada_sudoku/
RUN chmod +x /opt/tada_sudoku/bin/tada_sudoku

# Create appuser
RUN adduser -D -g '' tada_sudoku
USER tada_sudoku

ENTRYPOINT ["/usr/bin/dumb-init", "--"]

CMD ["/opt/tada_sudoku/bin/tada_sudoku < /opt/tada_sudoku/solved.txt"]
