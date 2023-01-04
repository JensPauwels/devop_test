FROM alpine:latest

WORKDIR /app

COPY devops_test ./devops_test

EXPOSE 1337

CMD ["./devops_test"]

