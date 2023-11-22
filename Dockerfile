FROM golang:latest

WORKDIR /app

COPY . .

RUN go mod init github.com/dlworhd/email-server

RUN go build -o email-server .

CMD ["./email-server"]