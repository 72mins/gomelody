FROM golang:1.22-alpine3.19

RUN apk update & apk upgrade
RUN apk add build-base pkgconfig git opus-dev
RUN apk --no-cache add ca-certificates wget bash xz-libs

WORKDIR /tmp

RUN wget https://johnvansickle.com/ffmpeg/releases/ffmpeg-release-amd64-static.tar.xz
RUN tar -xvJf  ffmpeg-release-amd64-static.tar.xz
RUN cd ff* && mv ff* /usr/local/bin

WORKDIR /app

COPY go.mod go.sum ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=1 go build -o main .

EXPOSE 8080

CMD ["./main"]
