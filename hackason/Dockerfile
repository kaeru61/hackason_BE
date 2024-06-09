FROM golang:1.22 as build
WORKDIR /hackason
COPY . /hackason/
RUN go build -o main
EXPOSE 8000
CMD ["go", "run", "main.go"]