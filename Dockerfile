FROM golang:1.18

RUN go version
ENV GOPATH=/

COPY ./ ./
COPY ./schema/000001_init.up.sql /docker-entrypoint-initdb.d/init.sql

RUN go mod download
RUN go build -o Run_HSE_Run ./cmd/main.go

CMD ["./Run_HSE_Run"]