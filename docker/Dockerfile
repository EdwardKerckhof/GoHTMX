# BUILD
FROM golang:1.21-alpine3.18 AS build

WORKDIR /apps/gohtmx/api

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build

# RUN
FROM alpine:3.18
WORKDIR /apps/gohtmx/api

COPY --from=build /apps/gohtmx/api/api .

EXPOSE 1337
CMD ["./api"]

HEALTHCHECK --interval=20s --timeout=3s --retries=3 --start-period=60s CMD curl -f http://localhost:1337/healthcheck || exit 1