FROM alpine:3.19.1

RUN addgroup -S nonroot \
    && adduser -S nonroot -G nonroot

USER nonroot

WORKDIR /app

COPY dify-connector /app/dify-connector

CMD ["/app/dify-connector", "serve"]