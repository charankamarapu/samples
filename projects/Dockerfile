# Stage 1: Build the Golang application
FROM golang:1.18 AS go-builder
WORKDIR /app
COPY goapp/. ./
RUN go mod download

# Stage 2: Build the Node.js application
FROM node:14 AS js-builder
WORKDIR /app
COPY javaScript/. ./
RUN npm install

# Stage 3: Set up the Python environment
FROM ubuntu:latest as python-builder
RUN apt-get update && apt-get install -y python3 python3-pip curl lsof && apt-get clean && rm -rf /var/lib/apt/lists/*
WORKDIR /app
COPY python/. ./
RUN if [ -f "requirements.txt" ]; then pip install --no-cache-dir -r requirements.txt; fi

# install keploy
RUN curl --silent --location "https://github.com/keploy/keploy/releases/latest/download/keploy_linux_arm64.tar.gz" | tar xz -C /tmp
RUN mkdir -p /usr/local/bin && mv /tmp/keploy /usr/local/bin && chmod +x /usr/local/bin/keploy

# Final Stage: Assemble the final image
FROM ubuntu:latest as runtime
RUN apt-get update && apt-get install -y curl lsof sudo nano git golang-go

# Install Node.js and npm
RUN curl -sL https://deb.nodesource.com/setup_14.x | bash - && apt-get install -y nodejs

# Install Python and pip
RUN apt-get install -y python3 python3-pip

# Clean up to reduce image size
RUN apt-get clean && rm -rf /var/lib/apt/lists/*

# Set up the working directory
WORKDIR /app

# Copy the Golang application source code and dependencies
COPY --from=go-builder /app /app/go-app
COPY --from=go-builder /go/pkg/mod /go/pkg/mod
COPY --from=go-builder /go/pkg/mod /root/go/pkg/mod


# Copy the Node.js application code and node_modules
COPY --from=js-builder /app /app/js-app

# Copy the Python application code
COPY --from=python-builder /app /app/python-app
COPY --from=python-builder /usr/local/lib/python3.10/dist-packages /usr/local/lib/python3.10/dist-packages

COPY --from=python-builder /usr/local/bin/keploy /usr/local/bin/keploy
RUN chmod +x /usr/local/bin/keploy

ENV PYTHONUNBUFFERED 1

CMD ["sleep", "infinity"]
