FROM golang:alpine as builder
RUN mkdir /build
ADD . /build/
WORKDIR /build/cmd/mindesk
RUN go build -o main .

FROM alpine
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /build/cmd/mindesk/main /app/
WORKDIR /app
EXPOSE 8080
CMD ["./main"]