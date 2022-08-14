FROM golang:1.18.5-alpine3.16 AS build
WORKDIR /app
COPY . .
RUN go build -o dstrestart

# ===

FROM alpine:3.16
MAINTAINER Jaehyeon Park <skystar@skystar.dev>
COPY --from=build /app/dstrestart /
ENTRYPOINT ["/dstrestart"]
