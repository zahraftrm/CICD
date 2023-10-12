FROM golang:1.21.0-alpine

WORKDIR /yahuy

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /iniAPP

CMD ["/iniAPP"]
