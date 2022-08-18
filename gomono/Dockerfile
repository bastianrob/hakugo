# syntax=docker/dockerfile:1

FROM golang:1.19-buster AS build
ARG service

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /service cmd/${service}/main.go


FROM gcr.io/distroless/base-debian10
ARG port
ARG version
ENV PORT=${port}
ENV APP_VERSION=${version}

WORKDIR /

COPY --from=build /service /service

EXPOSE ${port}

USER nonroot:nonroot

ENTRYPOINT ["/service"]
