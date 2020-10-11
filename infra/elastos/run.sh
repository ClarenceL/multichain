#!/bin/bash
ADDRESS=$1

# Print setup
echo "ELASTOS_ADDRESS=$ADDRESS"

# We must run from this directory so it can import config.json
cd /root

# Run blockchain miner
./ela