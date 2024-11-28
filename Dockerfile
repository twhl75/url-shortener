FROM golang:1.23

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o /url-shortener ./cmd

EXPOSE 8080

CMD [ "/url-shortener" ]
