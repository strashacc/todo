from golang:latest

WORKDIR /project

COPY . .

RUN go mod download

ENV MONGO_URI="mongodb://analytics_db:27017"
ENV REST_PORT=8000

CMD [ "go", "run", "cmd/main.go" ]