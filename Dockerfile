# Configuration.
ARG GO_VERSION=1

# Builder image.
FROM golang:${GO_VERSION}-alpine AS builder
WORKDIR /src

  # Install dependencies (cached).
  COPY go.mod ./
  RUN go mod download && go mod verify

  # Build the binary.
  COPY . .
  RUN go build -v -o h24 .

# Runtime image.
FROM alpine
WORKDIR /fly

  # Static files
  COPY templates /fly/templates

  # Binary
  COPY --from=builder /src/h24 .

# Done.
CMD ["./h24"]
