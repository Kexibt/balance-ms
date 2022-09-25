FROM golang:latest
RUN mkdir /app
COPY . /app
WORKDIR /app
RUN go build -o service main.go 
CMD [ "/app/service" ]