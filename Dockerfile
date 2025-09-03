# -------- STAGE 1: Build --------
FROM golang:1.24.2-alpine AS builder
# Set working directory
WORKDIR /app
# Cache Go modules
COPY go.mod go.sum ./
RUN go mod download
# Copy the entire project
COPY . .
# Build the Go binary
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/server/main.go

# -------- STAGE 2: Runtime --------
FROM ubuntu:22.04
# Install Chrome and dependencies
RUN apt-get update && apt-get install -y \
    wget \
    gnupg \
    ca-certificates \
    && wget -q -O - https://dl.google.com/linux/linux_signing_key.pub | apt-key add - \
    && echo "deb [arch=amd64] http://dl.google.com/linux/chrome/deb/ stable main" > /etc/apt/sources.list.d/google-chrome.list \
    && apt-get update \
    && apt-get install -y google-chrome-stable \
    && rm -rf /var/lib/apt/lists/*

# Create a non-root user
RUN useradd -ms /bin/bash chrome

# Set working directory
WORKDIR /app
# Copy only the final binary from the builder stage
COPY --from=builder /app/main .
COPY --from=builder /app/emails/templates/* ./emails/templates/
COPY --from=builder /app/static/templates/* ./static/templates/

# Change ownership to chrome user
RUN chown -R chrome:chrome /app

# Switch to chrome user
USER chrome

# Expose the app port
EXPOSE ${PORT}
# Run the binary
CMD ["./main"]