# syntax=docker/dockerfile:1

FROM golang:1.24-alpine AS build

ARG VERSION=0.1.0-dev
ARG COMMIT=none
ARG DATE=unknown

WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 go build -trimpath \
    -ldflags "-s -w \
      -X github.com/be-thomas/openspl/pkg/version.Version=${VERSION} \
      -X github.com/be-thomas/openspl/pkg/version.Commit=${COMMIT} \
      -X github.com/be-thomas/openspl/pkg/version.Date=${DATE}" \
    -o /out/openspl ./cmd/openspl

FROM gcr.io/distroless/static-debian12:nonroot

COPY --from=build /out/openspl /openspl
USER nonroot:nonroot
ENTRYPOINT ["/openspl"]
