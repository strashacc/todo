FROM golang:latest

WORKDIR /project

COPY . .

RUN go mod download

ENV MONGO_URI="mongodb://team_db:27017"
ENV GRPC_PORT=50051
ENV DB=teams

CMD [ "go", "run", "main.go" ]