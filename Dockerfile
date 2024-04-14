FROM golang:alpine3.20 AS build

WORKDIR /app

COPY go.mod go.sum /app/
RUN go mod download
COPY ./ /app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app/connectr-api ./cmd/main.go

FROM gcr.io/distroless/static-debian12 AS release
COPY --from=build /app/connectr-api /app/dev.env /

EXPOSE 9000

ENTRYPOINT [ "/connectr-api" ]