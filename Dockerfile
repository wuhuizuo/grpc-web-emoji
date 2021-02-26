FROM golang:1 as builder
WORKDIR /server
COPY ./ .
RUN go build -o emoji-service

FROM alpine:3.13.2
WORKDIR /app

COPY --from=builder /server/emoji-service ./
ENTRYPOINT [ "./emoji-service" ]
CMD [ "9000" ]

EXPOSE 9000