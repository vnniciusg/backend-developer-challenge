FROM golang:alpine AS build

WORKDIR /app

COPY . .

RUN go test ./...

RUN go build -o main cmd/main.go

FROM scratch as prod

WORKDIR /opt/app

COPY --from=build /app/main ./
COPY --from=build /app/.env ./
COPY --from=build /app/database/migrations /opt/app/database/migrations

EXPOSE 8080

CMD ["./main"]
