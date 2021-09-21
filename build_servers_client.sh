#!/bin/bash
RED='\033[0;31m'
NC='\033[0m' # No Color
Cyan='\033[0;36m'

# docker build
docker build -t imageslip .

pushd client

# build client cli application
go build -o upload_image

popd

# info
echo -e "${RED}use upload_image <image_path> for uploading after running run_server_docker.sh ${NC}"
