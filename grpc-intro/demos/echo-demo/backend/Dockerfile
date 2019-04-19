
FROM golang:alpine as builder
RUN mkdir -p "/go/src/backend"
COPY . "/go/src/backend"
RUN CGO_ENABLED=0 go install backend/cmd/echod

FROM scratch
COPY --from=builder /go/bin/echod /echod
ENV PORT 9000
EXPOSE 9000
ENTRYPOINT ["/echod"]
