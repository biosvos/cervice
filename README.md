[go-micro](https://github.com/go-micro/go-micro) v4를 이용한 테스트
#### dependency
```bash
go get go-micro.dev/v4
```

#### proto 생성기 설치
```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install github.com/go-micro/generator/cmd/protoc-gen-micro@latest 
```
#### proto 생성
```bash
protoc --proto_path=. --micro_out=. --go_out=. proto/greeter.proto
```

### test
#### go-micro 설치
```bash
go install github.com/go-micro/cli/cmd/go-micro@latest
```

#### service 목록 확인
```bash
go-micro services
```
#### inspect service
```bash
SERVICE_NAME="cervice"
go-micro describe service ${SERVICE_NAME:?}
```
#### call method
```bash
SERVICE_NAME="cervice"
ENDPOINT_NAME="Handler.Hello"
go-micro call ${SERVICE_NAME:?} ${ENDPOINT_NAME:?} '{"name": "John"}'
```
