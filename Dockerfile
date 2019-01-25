FROM golang as builder
WORKDIR /go/src/github.com/graphql-services/graphql-event-store-changelog
COPY . .
RUN go get ./... 
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /tmp/app .

FROM alpine:3.5

WORKDIR /app

COPY --from=builder /tmp/app /usr/local/bin/app
COPY --from=builder /go/src/github.com/graphql-services/graphql-event-store-changelog/schema.graphql /app/schema.graphql

# RUN apk --update add docker

# https://serverfault.com/questions/772227/chmod-not-working-correctly-in-docker
RUN chmod +x /usr/local/bin/app

ENTRYPOINT []
CMD [ "app", "server" ]