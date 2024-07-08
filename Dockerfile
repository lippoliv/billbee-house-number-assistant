# Build binary
FROM --platform=$BUILDPLATFORM golang:1.22.5-alpine AS build-env

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

# Ensure up to date root certificates, SSL verification will fail without it
COPY --from=build-env /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# "Install" our binary
COPY --from=build-env /app/assistant /

ENTRYPOINT ["/assistant"]
