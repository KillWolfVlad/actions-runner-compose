################################################################################

ARG BUILDER_IMAGE=golang:1.25.0-alpine3.22
ARG RUNTIME_IMAGE=alpine:3.23.2

################################################################################

FROM ${BUILDER_IMAGE} AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o arc main.go

################################################################################

FROM ${RUNTIME_IMAGE} AS runtime

WORKDIR /app

COPY --from=builder /app/arc .

ENTRYPOINT ["/app/arc"]
