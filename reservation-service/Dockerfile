FROM golang:1.24-alpine AS build
WORKDIR /src
COPY . .
RUN CGO_ENABLED=0 go build -o /reservation .
FROM scratch
WORKDIR /app
COPY --from=build /reservation .
COPY swagger ./swagger
ENTRYPOINT ["./reservation"]
