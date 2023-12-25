FROM golang:1.21-alpine

WORKDIR /app_tg
COPY go.mod go.sum ./
RUN go mod download

RUN mkdir "/app_tg/cmd"
COPY /cmd /app_tg/cmd
RUN mkdir "/app_tg/pkg"
COPY /pkg /app_tg/pkg

WORKDIR "/app_tg/cmd"
RUN go build -o main .

EXPOSE 8084
EXPOSE 8085

CMD ["./main"]