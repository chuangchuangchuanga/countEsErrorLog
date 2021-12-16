FROM golang:1.16.5-alpine as BUILD

COPY . /usr/app/

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
	GOPROXY="https://goproxy.cn,direct"

WORKDIR /usr/app/
RUN go build -o app .


FROM golang:1.16.5-alpine


RUN mkdir /app

WORKDIR /app

COPY --from=BUILD /usr/app/app /app/

EXPOSE 2112

CMD ["/app/app"]