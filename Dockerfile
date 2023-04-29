# Simple usage with a mounted data directory:
# > docker build -t gridchain .
# > docker run -it -p 36657:36657 -p 36656:36656 -v ~/.gridchaind:/root/.gridchaind -v ~/.gridchaincli:/root/.gridchaincli gridchain gridchaind init mynode
# > docker run -it -p 36657:36657 -p 36656:36656 -v ~/.gridchaind:/root/.gridchaind -v ~/.gridchaincli:/root/.gridchaincli gridchain gridchaind start
FROM golang:1.17.2-alpine AS build-env

# Install minimum necessary dependencies, remove packages
RUN apk add --no-cache curl make git libc-dev bash gcc linux-headers eudev-dev

# Set working directory for the build
WORKDIR /go/src/github.com/gridironx/gridchain

# Add source files
COPY . .

ENV GO111MODULE=on \
    GOPROXY=http://goproxy.cn
# Build GRIDIronxChain
RUN make install

# Final image
FROM alpine:edge

WORKDIR /root

# Copy over binaries from the build-env
COPY --from=build-env /go/bin/gridchaind /usr/bin/gridchaind
COPY --from=build-env /go/bin/gridchaincli /usr/bin/gridchaincli

# Run gridchaind by default, omit entrypoint to ease using container with gridchaincli
CMD ["gridchaind"]
