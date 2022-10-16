FROM golang:1.19-alpine AS builder
WORKDIR /app
COPY . ./
RUN apk update && apk add --no-cache git && apk add --update npm
RUN npm --prefix web install
RUN npm --prefix web run build
RUN rm -rf web/node_modules
RUN go mod download
RUN CGO_ENABLED=0  go build -o /usr/bin/dnsresolver ./cmd/dnsresolver
############################### SECOND STAGE ###############################
FROM alpine
COPY --from=builder /app/web /web
COPY --from=builder /usr/bin/dnsresolver /usr/bin/dnsresolver
EXPOSE 80
CMD ["dnsresolver"]
