FROM golang

WORKDIR /usr/app

COPY . .

RUN go build

CMD ./api-with-rate-limiter