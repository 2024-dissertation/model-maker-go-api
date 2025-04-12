FROM ubuntu:22.04 AS builder

ENV TZ=Europe/Minsk

RUN ln -snf /usr/share/zoneinfo/$CONTAINER_TIMEZONE /etc/localtime && echo $CONTAINER_TIMEZONE > /etc/timezone

RUN apt-get update && DEBIAN_FRONTEND=noninteractive TZ=Etc/UTC apt-get install -y --no-install-recommends \
        tzdata \
    && rm -rf /var/lib/apt/lists/*

RUN apt-get update && DEBIAN_FRONTEND=noninteractive TZ=Etc/UTC apt-get install -y \
  cmake \
  build-essential \
  graphviz \
  git \
  coinor-libclp-dev \
  libceres-dev \
  libjpeg-dev \
  liblemon-dev \
  libpng-dev \
  libtiff-dev \
  wget \
  libglu1-mesa-dev \
  libglew-dev \
  libglfw3-dev \
  python3-dev \
  libboost-system-dev \
  libboost-thread-dev \
  libboost-program-options-dev \
  libboost-test-dev \
  libboost-iostreams-dev libboost-program-options-dev libboost-system-dev \
  libboost-serialization-dev libboost-filesystem-dev libboost-thread-dev \
  libboost-regex-dev libboost-graph-dev \
  libopencv-dev \
  libcgal-dev \
  libcgal-qt5-dev \
  python3; \
  apt-get autoclean && apt-get clean

RUN git clone https://github.com/cdcseacave/VCG.git vcglib

# Install Go (optional if FROM golang:latest in next stage)
ENV GOLANG_VERSION=1.21.5
RUN wget https://go.dev/dl/go${GOLANG_VERSION}.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go${GOLANG_VERSION}.linux-amd64.tar.gz && \
    rm go${GOLANG_VERSION}.linux-amd64.tar.gz
ENV PATH="/usr/local/go/bin:$PATH"

# Copy app source code
WORKDIR /app
COPY . .

# --- Build OpenMVG ---
RUN mkdir /app/cmd/openMVG_build && cd /app/cmd/openMVG_build && \
    cmake -DCMAKE_BUILD_TYPE=RELEASE /app/cmd/openMVG/src && \
    make -j$(nproc)

# --- Build OpenMVS ---
RUN mkdir /app/cmd/openMVS_build && cd /app/cmd/openMVS_build && \
    cmake -DCMAKE_BUILD_TYPE=RELEASE \
          -DOpenMVG_DIR=/app/cmd/openMVG_build/src/openMVG/cmake/ \
          -DVCG_ROOT=/vcglib \
          /app/cmd/openMVS && \
    make -j$(nproc)

# --- Build Go App ---
ENV GOPATH=/go
ENV PATH="$PATH:/go/bin"
RUN go mod vendor
RUN go mod download
RUN go install github.com/pressly/goose/v3/cmd/goose@latest
RUN go build -o /app/main .

# --- Final image ---
FROM ubuntu:22.04

# Install runtime dependencies
RUN apt-get update && DEBIAN_FRONTEND=noninteractive TZ=Etc/UTC apt-get install -y \
    libpng-dev libjpeg-dev libtiff-dev \
    libglfw3-dev libglew-dev \
    libboost-iostreams-dev libboost-program-options-dev libboost-system-dev \
    libboost-serialization-dev libboost-filesystem-dev libboost-thread-dev \
    libboost-regex-dev libboost-graph-dev \
    libatlas-base-dev libopencv-dev \
    libprotobuf-dev protobuf-compiler \
    libgoogle-glog-dev libgflags-dev \
    && rm -rf /var/lib/apt/lists/*

# Copy built Go app and OpenMVG/OpenMVS binaries
COPY --from=builder /app/main /usr/local/bin/main
COPY --from=builder /app/cmd/openMVG_build/Linux-x86_64-RELEASE /usr/local/bin/openMVG
COPY --from=builder /app/cmd/openMVS_build/bin /usr/local/bin/openMVS

ENV PATH="/usr/local/bin/openMVG:/usr/local/bin/openMVS:$PATH"

WORKDIR /app
# CMD ["main"]

EXPOSE 3333