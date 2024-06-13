FROM --platform=linux/amd64 golang:1.22 as build
WORKDIR /db
COPY . /db/
RUN go build -o main
EXPOSE 8080
CMD ["go", "run", "main.go"]