module github.com/tomazis/kioku/server/srv-session-api

go 1.18

require (
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	github.com/snovichkov/zap-gelf v1.1.0
	github.com/tomazis/kioku/server/srv-session-api/pkg/srv-session-api v0.0.0-00010101000000-000000000000
	go.uber.org/zap v1.21.0
	google.golang.org/grpc v1.45.0
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
)

require (
	github.com/envoyproxy/protoc-gen-validate v0.6.7 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.10.0 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.8.0 // indirect
	golang.org/x/net v0.0.0-20220127200216-cd36cc0744dd // indirect
	golang.org/x/sys v0.0.0-20211216021012-1d35b9e2eb4e // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20220414192740-2d67ff6cf2b4 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
)

replace github.com/tomazis/kioku/server/srv-session-api/pkg/srv-session-api => ./pkg/srv-session-api
