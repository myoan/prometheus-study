FROM golang:1.17.8

WORKDIR /app
COPY . /app
RUN go build -o app
CMD ./app
