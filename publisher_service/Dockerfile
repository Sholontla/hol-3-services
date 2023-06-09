FROM golang:1.19

WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download

# Install gcloud CLI
RUN curl -O https://dl.google.com/dl/cloudsdk/channels/rapid/downloads/google-cloud-sdk-434.0.0-linux-x86_64.tar.gz && \
    tar -xf google-cloud-sdk-434.0.0-linux-x86_64.tar.gz && \
    ./google-cloud-sdk/install.sh --quiet && \
    rm google-cloud-sdk-434.0.0-linux-x86_64.tar.gz

# Set GOOGLE_APPLICATION_CREDENTIALS environment variable
ENV GOOGLE_APPLICATION_CREDENTIALS=/app/config/keys/hol-3-service-12288c028bbc.json

COPY ${ADMIN_DIR} .

CMD ["go", "run", "/app/cmd/main.go"]