FROM golang:alpine as pre-build

WORKDIR /go/src/dingbot
COPY . .
RUN go build -o dingbot cmd/main.go

FROM alpine
WORKDIR /app/

ENV TZ=Asia/Shanghai

COPY --from=pre-build /go/src/dingbot/dingbot /usr/local/bin
RUN echo "http://mirrors.aliyun.com/alpine/v3.10/main/" > /etc/apk/repositories && \
    apk --update --no-cache add ca-certificates