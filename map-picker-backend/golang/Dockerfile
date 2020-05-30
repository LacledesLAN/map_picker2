FROM golang:1.14.3-alpine AS builder

WORKDIR /build

# Copy Code and download dependency using go mod
COPY . .
RUN go mod download

# Build the App
RUN CGO_ENABLED=0 go build -o /dist/main ./cmd/mappicker



############################
# STEP 2 build a small image
############################
FROM scratch

COPY --from=builder /dist/main /
COPY ./data-responses/ /data-responses/

EXPOSE 8080

ENTRYPOINT [ "/main" ]
