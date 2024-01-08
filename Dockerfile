FROM debian:12.4-slim

MAINTAINER Oliver Lippert <oliver@lipperts-web.de>

ENV BILLBEE_USER=""
ENV BILLBEE_PASSWORD=""
ENV BILLBEE_API_KEY=""

COPY src/assistant /assistant
RUN chmod +x /assistant

ENTRYPOINT ["/assistant"]
