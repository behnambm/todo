FROM golang:1.20
LABEL authors="behnam"

RUN mkdir /app
WORKDIR /app

COPY ./bin/app.out .
COPY ./bin/db-init.out .

# Initialize the database
RUN ["/app/db-init.out", "--initdb"]

CMD ["/app/app.out"]

