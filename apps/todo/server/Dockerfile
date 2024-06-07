FROM golang:1.21 AS build

WORKDIR /demo

COPY ./apps/todo/server/src/go.mod ./apps/todo/server/src/go.mod
COPY ./apps/todo/server/src/go.sum ./apps/todo/server/src/go.sum
COPY ./libs/go ./libs/go
RUN go work init
RUN go work use -r .
RUN go mod download

COPY ./apps/todo/server/src ./apps/todo/server/src

WORKDIR /demo/apps/todo/server
RUN go build -o /bin/server ./src

FROM debian:12

RUN apt update -y && \
    apt upgrade -y && \
    apt install -y ca-certificates

WORKDIR /

COPY ./apps/todo/docs/openapi.yaml /var/openapi.yaml
COPY --from=build bin/server bin/server

EXPOSE 3000

ENTRYPOINT ["/bin/server"]
