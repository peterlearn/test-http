FROM golang:1.20 AS builder

WORKDIR /src

COPY . .

WORKDIR /src/activity

RUN git config --global url."https://oauth2:glpat-iehtnh3GhsW-4Gp3eAyo@gitlab.com".insteadOf "https://gitlab.com" \
    && go env -w GOPRIVATE=gitlab.com \
    && go env -w GOPROXY=https://goproxy.cn,direct && go env -w GOSUMDB="sum.golang.google.cn" \
    && go mod edit -replace test=github.com/peterlearn/test-proto@master
# To workaround GOPROXY unexpected EOF we run it a second time to repull failed packages
RUN --mount=type=cache,target=/root/.cache/go-build go mod tidy || go mod tidy 
RUN --mount=type=cache,target=/root/.cache/go-build make build

FROM debian:stable-slim

RUN chmod 777 /tmp && apt-get update && apt-get install -y --no-install-recommends \
		ca-certificates  \
        netbase \
        && rm -rf /var/lib/apt/lists/ \
        && apt-get autoremove -y && apt-get autoclean -y \
        && ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

RUN mkdir -p /opt/service/bin
WORKDIR /opt/service

COPY --from=builder /src/activity/bin  /opt/service/bin/
COPY --from=builder /src/activity/configs/srv_cfg_local /opt/service/configs/base

ENTRYPOINT ["/opt/service/bin/cmd", "-conf", "/opt/service/configs/base"]
