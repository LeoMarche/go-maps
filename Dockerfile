FROM golang:1.17-alpine

# Prepare
RUN adduser \
    --disabled-password \
    --no-create-home \
    --uid 1001 \
    appuser
RUN apk add build-base
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .

# Build
RUN go build -o /gomaps

# Run as user 1001
USER 1001
EXPOSE 8080
ENTRYPOINT ["/app/gomaps"]

# TODO : make multi-stage to reduce final image size (use alpine vanilla ? keep user 1001 ? copy the db ?)