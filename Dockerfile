FROM golang:alpine

# RUN apk add --no-cache git

WORKDIR /app

COPY . .

CMD ["go", "run", "main.go"]

EXPOSE 8081
