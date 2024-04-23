FROM golang:1.20

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64


WORKDIR /build


COPY go.mod .
COPY go.sum .
RUN go mod download


COPY . .


RUN go build -o main .


WORKDIR /dist
RUN rm -rf /etc/localtime
RUN ln -s /usr/share/zoneinfo/Asia/Jakarta /etc/localtime


RUN cp /build/main .
COPY .env /dist

EXPOSE 8787


CMD /dist/main consumer -e=$env