ARG GOLANG_VERSION=1.23.2

FROM golang:${GOLANG_VERSION}-alpine
ENV CGO_ENABLED 0

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh build-base curl

RUN curl -sS https://webinstall.dev/watchexec | bash
ENV PATH="${PATH}:/root/.local/bin"

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
ADD . .
EXPOSE 3333
CMD ["./auto.sh"]