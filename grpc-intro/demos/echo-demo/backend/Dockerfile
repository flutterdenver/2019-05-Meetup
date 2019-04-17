
FROM golang:alpine as builder
RUN mkdir -p "/go/src/backend"
COPY . "/go/src/backend"
RUN CGO_ENABLED=0 go install backend

FROM scratch
COPY --from=builder /go/bin/backend /backend
ENTRYPOINT ["/backend"]
