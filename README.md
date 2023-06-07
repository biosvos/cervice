[go-micro](https://github.com/go-micro/go-micro)를 이용한 테스트
- v4 기준

```bash
go get go-micro.dev/v4
```

- proto
```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install github.com/go-micro/generator/cmd/protoc-gen-micro@latest 
```
```bash
protoc --proto_path=. --micro_out=. --go_out=. proto/greeter.proto
```

# test
```bash
go-micro services
```

```bash
go install github.com/go-micro/cli/cmd/go-micro@latest
go-micro call helloworld Helloworld.Call '{"name": "John"}'
```

```bash
go-micro describe service cervice
```

```bash
go-micro call cervice Handler.Hello '{"name": "John"}'
```