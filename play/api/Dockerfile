FROM golang:1.20-buster AS build
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . ./
ENV GOARCH=amd64
RUN go build -o /go/bin/app

## Deploy
FROM gcr.io/distroless/base-debian11
COPY --from=build /go/bin/app /app
ENV PORT=8080
EXPOSE 8080
USER nonroot:nonroot
CMD ["/app"]