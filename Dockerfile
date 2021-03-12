FROM golang:1.16-alpine as builder

RUN apk update && apk add --no-cache make git

WORKDIR /app
COPY . /app

RUN CGO_ENABLED=0 make build

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /app/.config .config
COPY --from=builder /app/{MODULE} .

CMD ["./{MODULE}"]