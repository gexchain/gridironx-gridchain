FROM golang:stretch as build-env

# Install minimum necessary dependencies
ENV PACKAGES curl make git libc-dev bash gcc
RUN apt-get update && apt-get upgrade -y && \
    apt-get install -y $PACKAGES

WORKDIR /gridchain
# Add source files
COPY . .

# build gridchain
RUN make build-linux

# Final image
FROM golang:1.17 as final

WORKDIR /gridchaind
# Copy over binaries from the build-env
COPY --from=build-env /gridchain/build/gridchaind /usr/bin/gridchaind
COPY --from=build-env /gridchain/build/gridchaincli /usr/bin/gridchaincli
COPY --from=build-env /gridchain/networks/local/node/wrapper.sh /usr/bin/wrapper.sh

EXPOSE 26656 26657
ENTRYPOINT ["/usr/bin/wrapper.sh"]
CMD ["start"]
STOPSIGNAL SIGTERM

