FROM golang:latest

RUN go install github.com/cosmtrek/air@latest

WORKDIR /app

COPY . /app

EXPOSE 8080

CMD ["air", "-c", ".air.toml"]
