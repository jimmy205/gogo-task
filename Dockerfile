# build stage
FROM golang:1.18-rc-alpine AS build-env
WORKDIR /src

RUN apk add --no-cache
COPY . .
RUN go build -o server ./command/*.go

# final stage
FROM golang:1.18-rc-alpine
WORKDIR /app
COPY --from=build-env /src/server server
ENTRYPOINT ./server



