FROM golang:1.18

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o Run_HSE_Run ./cmd/main.go

EXPOSE 8080
EXPOSE 587

CMD ["./Run_HSE_Run"]