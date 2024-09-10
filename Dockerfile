FROM golang:1.22.2

RUN apt-get update && apt-get install postgresql-client make -y
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

WORKDIR /app

COPY . .
RUN go mod vendor
RUN go build -mod vendor -o /usr/local/bin/app

ENTRYPOINT [ "app" ]