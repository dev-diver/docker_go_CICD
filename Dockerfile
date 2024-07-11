# 베이스 이미지 설정 (Go 1.16)
FROM golang:1.16-alpine

# 작업 디렉토리 설정
WORKDIR /app

# 필요한 패키지 설치
COPY go.mod ./
COPY go.sum ./
RUN go mod tidy

# 소스 코드 복사
COPY *.go ./

# 애플리케이션 빌드
RUN go build -o /echo-app

# 컨테이너 실행
CMD ["/echo-app"]