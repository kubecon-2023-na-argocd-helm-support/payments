FROM golang:1.21 as builder
WORKDIR /src
ADD ./go.mod ./go.sum ./
RUN go mod download
ADD . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/payments ./main.go

FROM scratch
COPY --from=builder /bin/payments /bin/payments