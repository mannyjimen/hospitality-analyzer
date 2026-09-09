# syntax=docker/dockerfile:1
FROM golang:1.24 AS build-stage

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /hospitality

#deploying into binary lean image
FROM gcr.io/distroless/base-debian11 AS release-stage

WORKDIR /

COPY --from=build-stage /hospitality /hospitality

USER nonroot:nonroot

ENTRYPOINT ["/hospitality"]