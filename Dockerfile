FROM golang:1.13.5 as build

ENV CGO_ENABLED 0
ENV GO111MODULE on
ENV GOOS linux

WORKDIR /go/src/app
COPY . .

RUN go build -a -installsuffix cgo -o app ./cmd/gateway/main.go

# ---
FROM alpine:3.11

# Add non root user and certs
RUN apk --no-cache add ca-certificates \
  && addgroup -S app && adduser -S -g app app \
  && mkdir -p /home/app \
  && chown app /home/app

WORKDIR /home/app

COPY --from=build /go/src/app/app .

RUN chown -R app /home/app

USER app

EXPOSE 8080
CMD ["./app"]