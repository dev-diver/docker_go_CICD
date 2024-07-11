#!/bin/bash

# 도커 허브에서 최신 이미지 받기
docker pull devdiver/riddlefox:latest

# 기존 컨테이너 중지 및 삭제
docker stop riddlefox || true
docker rm riddlefox || true

# 새 컨테이너 실행
docker run -d --name riddlefox -p 5001:5001 devdiver/riddlefox:latest