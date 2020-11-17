# gin에서 swagger 설치하기

## gin 설치
```
$ go get -u github.com/gin-gonic/gin
```

## swag 설치
```
$ go get -u github.com/swaggo/swag/cmd/swag
$ mkdir swagger
$ touch main.go
$ swag init
```

## swagger 설치
```
$ go get -u github.com/swaggo/gin-swagger
```

## swag & swagger 연동
```
$ go get -u github.com/swaggo/gin-swagger/swaggerFiles
```

소스 코드 작성 후 swagger 코드로 변경하기 위해 `swag init` 명령어 실행
```
$ go run main.go
```

서버 실행 후 페이지 확인

http://localhost:8080/swagger/index.html
