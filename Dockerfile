FROM golang:1.18-buster as builder

WORKDIR /go/src/github.com/chiefcake/ergoproxy

COPY go.* .

RUN go mod download

COPY . .

RUN go build -o ergoproxy main.go

FROM alpine:3.16.0

RUN apk --no-cache add ca-certificates

COPY --from=builder /go/src/github.com/chiefcake/ergoproxy .

CMD [ "./ergoproxy" ]
