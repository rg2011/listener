FROM golang:1.14 AS builder

RUN mkdir /app
COPY . /app

WORKDIR /app
RUN go get
RUN CGO_ENABLED=0 go build -o /app/listener

FROM scratch
COPY --from=builder /app/listener /listener

EXPOSE 3000
ENTRYPOINT ["/listener"]
