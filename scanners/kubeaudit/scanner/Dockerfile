FROM golang:1.15.1 AS builder

# no need to include cgo bindings
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

# this is where we build our app
WORKDIR /go/src/app/

RUN git clone https://github.com/Shopify/kubeaudit.git /go/src/app/
RUN go mod download

RUN go build -a -ldflags '-w -s -extldflags "-static"' -o /go/bin/kubeaudit ./cmd/ \
  && chmod +x /go/bin/kubeaudit

FROM alpine:3.12
COPY --from=builder /go/bin/kubeaudit /kubeaudit
COPY wrapper.sh /wrapper.sh
RUN addgroup --system --gid 1001 kubeaudit && adduser kubeaudit --system --uid 1001 --ingroup kubeaudit
USER 1001
ENTRYPOINT ["/kubeaudit"]
CMD ["all"]
