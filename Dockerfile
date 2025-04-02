FROM golang:1.24-bookworm AS base

RUN apt-get update && \
    apt-get install -y --no-install-recommends make && \
    apt-get clean && rm -rf /var/lib/apt/lists/*

FROM base AS install

WORKDIR /app

COPY ./go.mod .
COPY ./go.sum .

COPY . .
RUN go mod download

RUN make build

FROM debian:bookworm-slim AS final

RUN apt-get update && \
    apt-get install -y --no-install-recommends ca-certificates tzdata jq curl && \
    apt-get clean && rm -rf /var/lib/apt/lists/*

COPY --from=install /app/build/nats-client ./nats-client

RUN groupadd -r app-user && useradd -r -g app-user app-user

USER app-user

CMD [ "./nats-client", "subscriber", "-s", "*" ]
