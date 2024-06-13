FROM --platform=linux/amd64 golang:1.22 as build
WORKDIR /hackason
COPY . /hackason/
RUN go build -o main
EXPOSE 8080
CMD ["go", "run", "main.go"]