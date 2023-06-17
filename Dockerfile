FROM golang:1.19.4 AS build

WORKDIR /app

COPY . ./

RUN go build -o /app/echo-server

FROM gcr.io/distroless/base-debian11:latest

COPY --from=build /app/echo-server /app/echo-server

USER nonroot

WORKDIR /app

ENTRYPOINT [ "./echo-server" ]