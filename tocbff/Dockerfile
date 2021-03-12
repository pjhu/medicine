FROM golang:alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build
COPY . .

RUN cd /build/application/main && go build -o app .

FROM scratch

COPY --from=builder /build/application/main/app /
COPY --from=builder /build/application/main/resources/ /resources/

ENTRYPOINT ["/app"]