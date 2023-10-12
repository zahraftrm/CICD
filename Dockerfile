FROM golang:1.21.0-alpine

WORKDIR /yahuy

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

# Ensure that the CGO_ENABLED=0 to build a static binary
RUN CGO_ENABLED=0 go build -o /iniAPP

CMD ["/iniAPP"]
