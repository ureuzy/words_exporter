FROM golang:1.18 as builder
WORKDIR /go/src/words_exporter
COPY . /go/src/words_exporter
RUN go get
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM alpine:latest
RUN apk --no-cache add tzdata && \
    cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime && \
    apk del tzdata
WORKDIR /work
COPY --from=builder /go/src/words_exporter/main /work
ENTRYPOINT ["./main"]
