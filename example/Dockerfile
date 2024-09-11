FROM golang:1.23 AS build-stage
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /bot main.go 

FROM gcr.io/distroless/static
WORKDIR /
COPY --from=build-stage /bot /bot
USER nonroot:nonroot
ENTRYPOINT ["/bot"]
