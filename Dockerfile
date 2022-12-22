FROM golang:1.18-alpine as builder

# Compile
WORKDIR /src/
COPY main.go go.* /src/
RUN CGO_ENABLED=0 go build -o /bin/omega-star

# Create a new scratch image
FROM scratch
COPY --from=builder /bin/omega-star /bin/omega-star
ENTRYPOINT ["/bin/omega-star"]