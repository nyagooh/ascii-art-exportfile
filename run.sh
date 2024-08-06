#!/bin/bash

docker build -t ascii-art-web-stylize .
docker run -d -p 8060:8060 --name ascii-art-web-stylize-container ascii-art-web-stylize
docker logs ascii-art-web-stylize-container
