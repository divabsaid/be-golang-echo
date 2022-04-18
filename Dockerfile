
# build
FROM golang:alpine as builder

RUN apk update && \
    apk add --no-cache git ca-certificates tzdata \
    # && apk add --no-cache curl \
    && update-ca-certificates

# Move to working directory /build
WORKDIR /be

# Copy go.mod & go.sum, run go mod download
COPY go.mod go.sum ./
RUN go mod download
# RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.6.2/migrate.linux-amd64.tar.gz | tar xvz
# RUN mv migrate.linux-amd64 /usr/bin/migrate

# Copy the code into the container
COPY . .

# Download dependency using go mod
# RUN go build -ldflags="-w -s" -o coklit .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o be .

# distribute
FROM alpine:3.13

WORKDIR /be

# Create appuser
ENV USER=appuser
ENV UID=10001
RUN adduser \    
    --disabled-password \    
    --gecos "" \    
    --home "/nonexistent" \    
    --shell "/sbin/nologin" \    
    --no-create-home \    
    --uid "${UID}" \    
    "${USER}"

COPY --from=builder --chown=appuser:appuser /be/be /be
# COPY --from=builder --chown=appuser:appuser coklit/db /coklit
# COPY --from=builder --chown=appuser:appuser /usr/bin/migrate /usr/bin/migrate
COPY config.yaml.example /be/config.yaml

USER appuser:appuser

STOPSIGNAL SIGINT

# Export necessary port
EXPOSE 8080

# Command to run when starting the container
ENTRYPOINT ["./be"]