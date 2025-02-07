# Start from the latest golang base image
FROM golang:1.23.1-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /workspace

# Copy the whole project
COPY . .

# Download all dependencies for the workspace
RUN go work sync

# Build the Go app
WORKDIR /workspace/app
#RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Start a new stage from scratch
FROM alpine:latest  

ENV TZ=Asia/Jakarta
RUN apk add --no-cache tzdata
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /workspace/app/main .

# Copy the .env file from the root of the project
COPY --from=builder /workspace/.env .

COPY --from=builder /workspace/static ./static

COPY --from=builder /workspace/shared/helper/firebase/keenos-notification-firebase-adminsdk-o19u7-eb419a0d95.json .

# Expose port 8081 to the outside world
EXPOSE 8081

# Command to run the executable
CMD ["./main"]