#!/bin/bash

# Configuration
REMOTE_HOST="104.248.96.212"
REMOTE_USER="root"
APP_NAME="perizinan"  # Service name
SERVICE_DIR="perizinan" # Service directory name
REMOTE_DIR="/opt/${APP_NAME}"
GO_BINARY="${APP_NAME}"
ENV_FILE=".env"
APP_PORT="8081"
BACKUP_DIR="${REMOTE_DIR}/backups"
UPLOADS_DIR="/var/www/${APP_NAME}/uploads"  # New uploads directory

# Colors for output
GREEN='\033[0;32m'
RED='\033[0;31m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Function to print status
print_status() {
    echo -e "${GREEN}[+] $1${NC}"
}

print_warning() {
    echo -e "${YELLOW}[!] $1${NC}"
}

print_error() {
    echo -e "${RED}[-] $1${NC}"
    exit 1
}

# Function to run remote commands
remote_exec() {
    ssh -o "BatchMode=yes" ${REMOTE_USER}@${REMOTE_HOST} "$1"
}

# Function to copy files
remote_copy() {
    scp -o "BatchMode=yes" "$1" ${REMOTE_USER}@${REMOTE_HOST}:"$2"
}

# Function to backup current deployment
backup_current_deployment() {
    print_status "Creating backup of current deployment..."
    local timestamp=$(date +%Y%m%d_%H%M%S)
    remote_exec "
        mkdir -p ${BACKUP_DIR}
        if [ -f ${REMOTE_DIR}/${GO_BINARY} ]; then
            cp ${REMOTE_DIR}/${GO_BINARY} ${BACKUP_DIR}/${GO_BINARY}_${timestamp}
        fi
        if [ -f ${REMOTE_DIR}/${ENV_FILE} ]; then
            cp ${REMOTE_DIR}/${ENV_FILE} ${BACKUP_DIR}/${ENV_FILE}_${timestamp}
        fi
    "
}

# Function to restore from backup
restore_from_backup() {
    print_warning "Deployment failed! Attempting to restore from backup..."
    local latest_backup=$(remote_exec "ls -t ${BACKUP_DIR}/${GO_BINARY}_* | head -n 1")
    local latest_env_backup=$(remote_exec "ls -t ${BACKUP_DIR}/${ENV_FILE}_* | head -n 1")
    
    if [ ! -z "$latest_backup" ]; then
        remote_exec "
            cp ${latest_backup} ${REMOTE_DIR}/${GO_BINARY}
            cp ${latest_env_backup} ${REMOTE_DIR}/${ENV_FILE}
            systemctl restart ${APP_NAME}
        "
        print_status "Restored from backup. Previous version should be running."
    else
        print_error "No backup found to restore from!"
    fi
}

# Function to check dependencies
check_dependencies() {
    print_status "Checking dependencies..."
    
    # Check if go is installed
    if ! command -v go &> /dev/null; then
        print_error "Go is not installed. Please install it from https://golang.org/doc/install"
    fi

    # Check if we're in the project root
    if [ ! -f "go.work" ]; then
        print_error "Please run this script from the project root directory (where go.work is located)"
    fi

    # Verify service directory exists
    if [ ! -d "${SERVICE_DIR}" ]; then
        print_error "Service directory '${SERVICE_DIR}' not found"
    fi

    # Check if .env file exists
    if [ ! -f "${ENV_FILE}" ]; then
        print_error "Environment file '${ENV_FILE}' not found in root directory"
    fi
}

# Function to setup remote requirements
setup_remote_requirements() {
    print_status "Setting up remote requirements..."
    if ! remote_exec "
        # Install required packages
        apt-get update
        apt-get install -y net-tools ufw
        
        # Create necessary directories
        mkdir -p ${REMOTE_DIR}
        mkdir -p ${BACKUP_DIR}
        mkdir -p ${UPLOADS_DIR}
        
        # Set proper permissions for uploads directory
        chmod 755 ${UPLOADS_DIR}
        chown ${REMOTE_USER}:${REMOTE_USER} ${UPLOADS_DIR}
        
        # Create .gitkeep to preserve directory
        touch ${UPLOADS_DIR}/.gitkeep
    "; then
        print_error "Failed to setup remote requirements"
    fi
}

# Function to handle the running service
handle_running_service() {
    print_status "Checking for running service..."
    PORT_CHECK=$(remote_exec "netstat -tuln | grep LISTEN | grep :${APP_PORT}")
    if [ ! -z "$PORT_CHECK" ]; then
        print_warning "Service is currently running on port ${APP_PORT}"
        remote_exec "systemctl status ${APP_NAME} --no-pager"
        read -p "Do you want to stop the service and continue deployment? (y/n) " -n 1 -r
        echo
        if [[ ! $REPLY =~ ^[Yy]$ ]]; then
            print_error "Deployment aborted"
        fi
        
        print_status "Stopping service..."
        remote_exec "systemctl stop ${APP_NAME}"
        sleep 3  # Wait for service to fully stop
    fi
}

# Function to build application
build_application() {
    print_status "Building ${APP_NAME}..."
    cd "${SERVICE_DIR}" || print_error "Failed to change to service directory"
    go clean
    GOOS=linux GOARCH=amd64 go build -o "${GO_BINARY}" ./main.go
    if [ $? -ne 0 ]; then
        print_error "Build failed"
    fi
    cd ..
}

# Function to deploy files
deploy_files() {
    print_status "Deploying files to remote server..."
    
    # Backup current deployment
    backup_current_deployment
    
    # Copy new binary
    if ! remote_copy "${SERVICE_DIR}/${GO_BINARY}" "${REMOTE_DIR}/"; then
        restore_from_backup
        print_error "Failed to copy binary to remote server"
    fi
    
    # Copy .env file
    if ! remote_copy "${ENV_FILE}" "${REMOTE_DIR}/"; then
        restore_from_backup
        print_error "Failed to copy .env file to remote server"
    fi
    
    # Set permissions
    remote_exec "
        chmod 755 ${REMOTE_DIR}/${GO_BINARY}
        chmod 600 ${REMOTE_DIR}/${ENV_FILE}
    "
}

# Function to setup systemd service
setup_systemd_service() {
    print_status "Setting up systemd service..."
    
    # Create service file
    cat << EOF > ${APP_NAME}.service
[Unit]
Description=${APP_NAME} service
After=network.target mariadb.service

[Service]
Type=simple
User=root
WorkingDirectory=${REMOTE_DIR}
Environment="UPLOAD_DIR=${UPLOADS_DIR}"
ExecStart=${REMOTE_DIR}/${GO_BINARY}
Restart=on-failure
RestartSec=5
StandardOutput=append:/var/log/${APP_NAME}.log
StandardError=append:/var/log/${APP_NAME}.log

[Install]
WantedBy=multi-user.target
EOF

    # Copy and enable service
    remote_copy "${APP_NAME}.service" "/etc/systemd/system/"
    remote_exec "
        systemctl daemon-reload
        systemctl enable ${APP_NAME}
        systemctl restart ${APP_NAME}
    "
}

# Function to verify deployment
verify_deployment() {
    print_status "Verifying deployment..."
    sleep 5  # Give application time to start
    
    # Check if service is running
    SERVICE_STATUS=$(remote_exec "systemctl is-active ${APP_NAME}")
    if [ "$SERVICE_STATUS" != "active" ]; then
        print_warning "Service failed to start. Checking logs..."
        remote_exec "
            systemctl status ${APP_NAME} --no-pager
            echo -e '\nApplication Logs:'
            tail -n 50 /var/log/${APP_NAME}.log
        "
        restore_from_backup
        print_error "Service failed to start properly"
    fi
    
    # Check if port is listening
    PORT_CHECK=$(remote_exec "netstat -tuln | grep LISTEN | grep :${APP_PORT}")
    if [ -z "$PORT_CHECK" ]; then
        print_warning "Application is not listening on port ${APP_PORT}"
        restore_from_backup
        print_error "Application failed to bind to port"
    fi
}

# Function to setup firewall
setup_firewall() {
    print_status "Configuring firewall..."
    remote_exec "
        ufw allow 22/tcp
        ufw allow ${APP_PORT}/tcp
        ufw --force enable
        ufw status
    "
}

# Function to cleanup
cleanup() {
    print_status "Cleaning up..."
    rm -f "${SERVICE_DIR}/${GO_BINARY}"
    rm -f "${APP_NAME}.service"
}

# Main deployment flow
main() {
    check_dependencies
    setup_remote_requirements
    handle_running_service
    build_application
    deploy_files
    setup_systemd_service
    verify_deployment
    setup_firewall
    cleanup
    
    print_status "Deployment completed successfully!"
    print_status "To check application logs, run:"
    echo "ssh ${REMOTE_USER}@${REMOTE_HOST} 'journalctl -u ${APP_NAME} -f'"
}

# Run main function
main