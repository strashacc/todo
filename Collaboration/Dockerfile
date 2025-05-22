FROM golang:latest

# RUN apt-get update
# RUN apt-get upgrade
# RUN apt-get install -y protobuf-compiler
# RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
# RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

WORKDIR /project

COPY . .

# RUN protoc --go_out=. --go-grpc_out=. proto/*
RUN go mod download

CMD [ "go", "run", "main.go" ]