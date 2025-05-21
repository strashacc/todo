FROM golang:latest

WORKDIR /project

COPY . .

RUN go mod download

CMD [ "go", "run", "main.go" ]