FROM golang
WORKDIR /app
COPY app.go /app
RUN go build app.go

CMD "/app/app"