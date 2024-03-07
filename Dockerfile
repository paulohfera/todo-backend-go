FROM golang:1.22.1-alpine3.19 as modules
COPY go.mod go.sum /modules/
WORKDIR /modules
RUN go mod download

FROM golang:1.22.1-alpine3.19 as builder
COPY --from=modules /go/pkg /go/pkg
COPY . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags migrate -o /bin/app main.go

FROM scratch
COPY --from=builder /bin/app /app
COPY --from=builder /app/migrations /migrations
CMD ["/app"]