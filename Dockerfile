FROM golang:1.23.3-bullseye

# Install TinyGo 0.34.0
RUN ARCH=$(dpkg --print-architecture) && \
    if [ "$ARCH" = "arm64" ]; then \
        wget https://github.com/tinygo-org/tinygo/releases/download/v0.34.0/tinygo_0.34.0_arm64.deb && \
        dpkg -i tinygo_0.34.0_arm64.deb && \
        rm tinygo_0.34.0_arm64.deb; \
    else \
        wget https://github.com/tinygo-org/tinygo/releases/download/v0.34.0/tinygo_0.34.0_amd64.deb && \
        dpkg -i tinygo_0.34.0_amd64.deb && \
        rm tinygo_0.34.0_amd64.deb; \
    fi

# Create workspace directory
WORKDIR /workspace

# Copy Makefile and install wescale_wasm
COPY Makefile .
RUN make install-wescale-wasm

# Add bin directory to PATH
ENV PATH="/workspace/bin:${PATH}"

# Copy entrypoint script
COPY <<'EOF' /entrypoint.sh
#!/bin/bash
set -e

# Build the WASM plugin
make build WASM_FILE=${WASM_FILE:-my_plugin.wasm}

# Deploy if DEPLOY is set to true
if [ "${DEPLOY}" = "true" ]; then
    wescale_wasm --command=install \
        --wasm_file=/workspace/bin/${WASM_FILE:-my_plugin.wasm} \
        --mysql_host=${MYSQL_HOST:-127.0.0.1} \
        --mysql_port=${MYSQL_PORT:-15306} \
        --mysql_user=${MYSQL_USER:-root} \
        --mysql_password=${MYSQL_PASSWORD:-root} \
        --create_filter \
        --filter_name=${FILTER_NAME:-my_plugin_wasm_filter} \
        --filter_desc="${FILTER_DESC:-Plugin created via Docker}" \
        --filter_status=${FILTER_STATUS:-ACTIVE} \
        --filter_plans="${FILTER_PLANS}"
fi

# Keep container running if KEEP_ALIVE is true
if [ "${KEEP_ALIVE}" = "true" ]; then
    tail -f /dev/null
fi
EOF

RUN chmod +x /entrypoint.sh

CMD ["/entrypoint.sh"]