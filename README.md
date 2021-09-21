# ImageSlip

A simple and plain image uploader.
    <br/>1. Upload image files
    <br/>2. Receive a URL after successful upload.
    <br/>3. Use/Share URL to download the image


## Build webserver and gRPCServer

Run ./build_servers_client.sh for building docker image. Both Webserver service
and  gRPCServer service runs on same container using supervisord

## Start servers

Run ./run_server_docker.sh for running docker image. 

## Run client

Run client as specified in build_servers_client.
./<built_binary> <path_to_image>
