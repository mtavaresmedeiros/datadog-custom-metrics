
##
## Build
##

FROM golang:1.18 AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /datadog-custom-metrics .

##
## Deploy
##

FROM gcr.io/distroless/base-debian11

WORKDIR /

COPY --from=build /datadog-custom-metrics /datadog-custom-metrics

USER nonroot:nonroot

ENTRYPOINT ["/datadog-custom-metrics"]