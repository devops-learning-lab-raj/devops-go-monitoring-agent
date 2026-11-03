FROM golang:1.23-alpine AS build
WORKDIR /src
COPY go.mod ./
COPY main.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /out/agent .
FROM scratch
COPY --from=build /out/agent /agent
EXPOSE 8080
ENTRYPOINT ["/agent"]