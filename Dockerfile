FROM golang:1.21.5-alpine

MAINTAINER Oliver Lippert <oliver@lipperts-web.de>

ENV BILLBEE_USER=""
ENV BILLBEE_PASSWORD=""
ENV BILLBEE_API_KEY=""

COPY src/assistant /assistant
RUN chmod +x /assistant

CMD /assistant
