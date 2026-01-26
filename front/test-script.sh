#!/usr/bin/env bash
set -euo pipefail

echo "Script started"
echo "Testing log function"

TRACE_COLOR=$'\033[1;34m'
RESET_COLOR=$'\033[0m'
log() { echo -e "${TRACE_COLOR}[$(date +%H:%M:%S)] $*${RESET_COLOR}"; }

log "Test message 1"
log "Test message 2"

echo "Script finished"
