FROM golang

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go mod tidy

EXPOSE 9001

CMD ["go","run","."]