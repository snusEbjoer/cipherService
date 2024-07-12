FROM golang:1.22.5

WORKDIR /cipherService/

COPY . .

RUN go mod download

WORKDIR /cipherService/cmd

RUN CGO_ENABLED=0 GOOS=linux go build -o ../bin

EXPOSE 42069

CMD [ "../bin/cmd" ]

