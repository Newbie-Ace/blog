FROM golang:1.20-buster

WORKDIR /app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN go build -v -o /app ./


EXPOSE  8000

CMD ["app"]