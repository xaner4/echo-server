FROM golang:1.24-alpine as builder

WORKDIR /srv/
COPY ./go.* ./
COPY ./*.go ./
RUN go mod vendor
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o echo-server *.go

FROM scratch
WORKDIR /srv/
COPY --from=builder /srv/echo-server /srv/echo-server

EXPOSE 4444

ENTRYPOINT [ "/srv/echo-server" ]
