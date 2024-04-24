# builder
FROM golang:alpine as builder

RUN apk update && apk add --no-cache git make

WORKDIR /app

COPY . .

RUN go mod tidy

RUN make build

# runner
FROM golang:alpine

RUN apk update 

WORKDIR /app

COPY --from=builder /app/build/core core

EXPOSE 8080

CMD ["./core"]
