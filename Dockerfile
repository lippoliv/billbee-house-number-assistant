# Build binary
FROM --platform=$BUILDPLATFORM golang:1.21.5-alpine AS build-env

ARG TARGETOS
ARG TARGETARCH

ADD src /app
WORKDIR /app

RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} \
    go build -ldflags="-s -w" -o assistant

# Build container
FROM scratch

MAINTAINER Oliver Lippert <oliver@lipperts-web.de>

ENV BILLBEE_USER=""
ENV BILLBEE_PASSWORD=""
ENV BILLBEE_API_KEY=""

COPY --from=build-env /app/assistant /

ENTRYPOINT ["/assistant"]

FROM debian:12.4-slim

MAINTAINER Oliver Lippert <oliver@lipperts-web.de>

ENV BILLBEE_USER=""
ENV BILLBEE_PASSWORD=""
ENV BILLBEE_API_KEY=""

COPY src/assistant /assistant
RUN chmod +x /assistant

ENTRYPOINT ["/assistant"]
