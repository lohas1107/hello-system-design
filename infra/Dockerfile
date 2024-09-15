FROM ubuntu:24.04 AS base
RUN apt-get update
RUN apt-get install -y make
RUN apt-get install -y git

FROM base AS go
ENV PATH="/usr/local/go/bin:${PATH}"
COPY --from=golang:latest /usr/local/go/ /usr/local/go/

FROM go AS workspace
WORKDIR /workspace