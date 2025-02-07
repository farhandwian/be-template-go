# Start from the latest golang base image
FROM golang:1.23.1-alpine AS builder
 
# Set the Current Working Directory inside the container
WORKDIR /workspace
 
# Copy the whole project
COPY . .
 
# Download all dependencies for the workspace
RUN go work sync
 
# Build binaries for each workspace
WORKDIR /workspace/dashboard
RUN CGO_ENABLED=0 GOOS=linux go build -o /workspace/bin/dashboard_main .
 
# Start a new stage from scratch
FROM alpine:latest
 
RUN apk --no-cache add ca-certificates
 
WORKDIR /root/
 
# Copy all binaries to the final image
COPY --from=builder /workspace/bin/dashboard_main ./dashboard_main
 
# Copy the .env file from the root of the project
COPY --from=builder /workspace/.env .
 
# Copy the seeder folder from the dashboard workspace to the root
COPY --from=builder /workspace/dashboard/seeder ./seeder
COPY --from=builder /workspace/static ./static
 
# Expose port 8081 to the outside world
EXPOSE 8081
 
# Default command to run the dashboard workspace
CMD ["./dashboard_main"]