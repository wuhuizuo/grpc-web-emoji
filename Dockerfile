FROM golang:1 as builder
WORKDIR /server
COPY ./ .
RUN go build -o emoji-service

FROM scratch
WORKDIR /bin/
COPY --from=builder /server/emoji-service /bin/emoji-service
ENTRYPOINT [ "/bin/emoji-service" ]
CMD [ "9000" ]

EXPOSE 9000