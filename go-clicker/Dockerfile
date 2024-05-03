FROM golang:1.22-alpine3.19 as build

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/go-clicker cmd/go-clicker/go-clicker.go

FROM alpine:3.19

COPY --from=build /usr/local/bin/go-clicker /usr/local/bin/go-clicker

CMD ["go-clicker"]
