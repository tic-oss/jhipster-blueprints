FROM golang

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN go mod tidy

EXPOSE <%= serverPort %>

CMD ["go","run","."]