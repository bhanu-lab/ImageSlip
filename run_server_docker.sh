#!/bin/bash
RED='\033[0;31m'
NC='\033[0m' # No Color
CYAN='\033[0;36m'

echo -e "${RED}Make sure latest docker image is build ${NC}"
echo -e "${CYAN}Wait for 10 seconds after server came up and before running client ${NC}"

# run docker container in detachmode with ports and image "imageslip"
sudo docker run -d -p 39298:39298 -p 8080:8080 -v imageslip:/tmp/ imageslip:latest
