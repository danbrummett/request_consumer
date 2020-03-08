FROM alpine

WORKDIR /app
ADD . ./
ENTRYPOINT ["./request_consumer"]
