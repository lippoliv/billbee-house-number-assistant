FROM golang:1.21.5-alpine

MAINTAINER Oliver Lippert <oliver@lipperts-web.de>

COPY src/assistant /assistant

CMD /assistant
