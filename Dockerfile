FROM golang:1.19.0 as base
ENV ENV=production
FROM base as dev
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
RUN go install github.com/swaggo/swag/cmd/swag@latest
WORKDIR /opt/app/api
COPY . .
RUN go mod download
CMD ["air"]