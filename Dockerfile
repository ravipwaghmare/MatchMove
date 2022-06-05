FROM golang:1.18.3 as builder

WORKDIR /workspace

COPY . .

# Build
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -ldflags " -extldflags -static" -o matchmove main.go

FROM alpine:latest

WORKDIR /

COPY --from=builder /workspace/matchmove .

COPY --from=builder /workspace/gorm.db .

EXPOSE 3000

ENTRYPOINT ["/matchmove"]
