FROM golang:1.18

WORKDIR /app

COPY . .
RUN go mod download

RUN go build -o /docker-go-ms .

ENV SERVER_PORT=8080
ENV DB_HOST=localhost
ENV DB_PORT=5432
ENV DB_NAME=user-ms-db
ENV DB_USER=user-ms
ENV DB_PASSWORD=secret

EXPOSE 8080

CMD [ "/docker-go-ms" ]