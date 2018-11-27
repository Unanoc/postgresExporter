FROM ubuntu:16.04
LABEL author="Daniel Lee"

# PostgreSQL installing
RUN apt-get -y update
RUN apt-get -y install apt-transport-https git wget
RUN echo 'deb http://apt.postgresql.org/pub/repos/apt/ xenial-pgdg main' >> /etc/apt/sources.list.d/pgdg.list
RUN wget --quiet -O - https://www.postgresql.org/media/keys/ACCC4CF8.asc | apt-key add -
RUN apt-get -y update
ENV PGVERSION 10
RUN apt-get -y install postgresql-$PGVERSION postgresql-contrib

# Golang installing
ENV GOVERSION 1.11.1
USER root
RUN wget https://storage.googleapis.com/golang/go$GOVERSION.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go$GOVERSION.linux-amd64.tar.gz && \
    mkdir go && mkdir go/src && mkdir go/bin && mkdir go/pkg
ENV GOROOT /usr/local/go
ENV GOPATH /opt/go
ENV PATH $GOROOT/bin:$GOPATH/bin:$PATH
RUN mkdir -p "$GOPATH/bin" "$GOPATH/src"
RUN apt-get -y install gcc musl-dev && GO11MODULE=on

# Downloading a project
WORKDIR /home
RUN mkdir postgresExporter
COPY . postgresExporter/
WORKDIR /home/postgresExporter/postgresExporter
RUN go build .
WORKDIR /home/postgresExporter/example
RUN go build .

# PostgreSQL creating of database
USER postgres
RUN /etc/init.d/postgresql start &&\
    psql --command "CREATE USER testing WITH SUPERUSER PASSWORD 'testing';" &&\
    createdb -O testing testing &&\
    psql -d testing -c "CREATE EXTENSION IF NOT EXISTS citext;" &&\
    psql testing -f schema.sql &&\
    ./example &&\
    /etc/init.d/postgresql stop

USER root
WORKDIR /home/postgresExporter/postgresExporter