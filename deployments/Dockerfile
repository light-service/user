# build stage
FROM golang:1.15.2-buster AS build-env

WORKDIR /root/user
ADD . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    && GO111MODULE="on" \
    && GOPROXY="https://goproxy.cn,direct" \
    && go mod download \
    && go test --cover  $(go list ./... | grep -v /vendor/) \
    && time go build -o main

# final stage
FROM golang:1.15.2-buster AS build-env
ENV TZ Asia/Shanghai

WORKDIR /app
RUN mkdir -p /app/ /app/migrations/ /data/
COPY --from=build-env /root/user/migrations /app/migrations
COPY --from=build-env /root/user/main /app/
COPY deploy/start.sh /app/
COPY deploy/clean_log.sh /app/
RUN chmod +x /app/main
RUN chmod +x /app/start.sh
RUN chmod +x /app/clean_log.sh
