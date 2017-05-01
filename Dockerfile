# Build stage
FROM golang:1.8 AS build-stage

RUN go version

RUN apt-get update
RUN curl -sL https://deb.nodesource.com/setup_7.x | bash
RUN apt-get install -y build-essential nodejs
RUN apt-get install -y sqlite3 libsqlite3-dev
RUN sh -c 'echo "deb http://apt.postgresql.org/pub/repos/apt/ `lsb_release -cs`-pgdg main" >> /etc/apt/sources.list.d/pgdg.list'
RUN wget -q https://www.postgresql.org/media/keys/ACCC4CF8.asc -O - | apt-key add -
RUN apt-get update
RUN apt-get install -y postgresql postgresql-contrib libpq-dev
RUN apt-get install -y -q mysql-client

RUN go get github.com/gobuffalo/buffalo/buffalo
RUN go get github.com/golang/dep/cmd/dep

ENV BP=$GOPATH/src/github.com/learnonline/learnonline

RUN mkdir -p $BP
WORKDIR $BP
ADD . .

RUN dep ensure
RUN buffalo version
RUN buffalo build


EXPOSE 3000

# Final Stage
FROM ubuntu:16.04

ARG GIT_COMMIT
ARG VERSION
LABEL REPO="https://github.com/learnonline/learnonline"
LABEL GIT_COMMIT=$GIT_COMMIT
LABEL VERSION=$VERSION

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:/opt/learnonline/bin

WORKDIR /opt/learnonline/bin

COPY --from=build-stage /go/src/github.com/learnonline/learnonline/bin/learnonline /opt/learnonline/bin/

COPY --from=build-stage /go/src/github.com/learnonline/learnonline/database.yml /opt/learnonline/bin/
COPY --from=build-stage /go/src/github.com/learnonline/learnonline/assets/ /opt/learnonline/assets

RUN chmod +x /opt/learnonline/bin/learnonline

CMD /opt/learnonline/bin/learnonline
