FROM golang:alpine AS build

RUN apk add make
RUN apk add gcc
RUN apk add libc-dev

RUN go install github.com/google/wire/cmd/wire@latest
RUN go install github.com/swaggo/swag/cmd/swag@latest

COPY . /app
WORKDIR /app

RUN go get
RUN make

FROM alpine

COPY --from=build /app/bin/server.app /bin/server.app
WORKDIR /app

RUN chmod +x /bin/server.app

CMD [ "server.app" ]
