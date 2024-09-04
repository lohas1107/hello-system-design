FROM ubuntu:24.04 AS base
RUN apt update
RUN apt install -y make

FROM base AS go

FROM go AS workspace
WORKDIR /workspace