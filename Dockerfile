FROM golang:1.22.3-bullseye AS build-base
WORKDIR /app
COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod \
  --mount=type=cache,target=/root/.cache/go-build \
  go mod download

RUN go install golang.org/x/vuln/cmd/govulncheck@latest
#------------------------------------------------
FROM golang:1.22.3-alpine AS dev
# Install air for hot reload & delve for debugging
WORKDIR /app
RUN go install github.com/cosmtrek/air@latest && \
  go install golang.org/x/vuln/cmd/govulncheck@latest && \
  go install github.com/go-delve/delve/cmd/dlv@latest && \
  go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest && \
  go install github.com/swaggo/swag/cmd/swag@latest
COPY go.mod go.sum ./
RUN go mod download
CMD ["air", "-c", ".air.toml"]

#------------------------------------------------
FROM build-base AS build-production
RUN useradd -u 1001 nonroot
COPY . .
RUN go build \
-ldflags="-linkmode external -extldflags -static" \
-tags netgo \
-o goapp
  
#------------------------------------------------
FROM alpine:3.20.0 AS production
ENV GIN_MODE=release
USER root
RUN adduser nonroot -D
WORKDIR /
# COPY --from=build-production /etc/passwd /etc/passwd
# COPY --from=build-production /app/healthcheck/healthcheck healthcheck
COPY --from=build-production --chown=nonroot:nonroot /app/goapp goapp
USER nonroot
EXPOSE 8080
CMD ["/goapp"]