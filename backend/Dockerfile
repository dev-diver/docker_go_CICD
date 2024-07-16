# 베이스 이미지 설정 (Go 1.16)
FROM golang:bookworm

RUN apt-get update && apt-get install -y \
  ca-certificates \
  curl \
  gnupg &&\
  install -m 0755 -d /etc/apt/keyrings && \
  curl -fsSL https://download.docker.com/linux/debian/gpg -o /etc/apt/keyrings/docker.asc && \
  chmod a+r /etc/apt/keyrings/docker.asc && \
  echo "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.asc] https://download.docker.com/linux/debian $(. /etc/os-release && echo "$VERSION_CODENAME") stable" | tee /etc/apt/sources.list.d/docker.list > /dev/null && \
  apt-get update && \
  apt-get install -y docker-ce docker-ce-cli && \
  rm -rf /var/lib/apt/lists/*

RUN groupadd -f docker && usermod -aG docker root

# 작업 디렉토리 설정
WORKDIR /app

# 필요한 패키지 설치
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# 소스 코드 복사
COPY *.go ./

# 애플리케이션 빌드
RUN go build -o /app/server

COPY docker-compose.yml /app/docker-compose.yml
COPY Dockerfile /app/Dockerfile

# 컨테이너 실행
CMD ["/app/server"]