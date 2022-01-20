module github.com/tomazis/kioku/server/srv-frontend-api

go 1.17

require (
	github.com/tomazis/kioku/server/srv-frontend-api/pkg/srv-frontend-api v0.0.0-00010101000000-000000000000
	go.uber.org/zap v1.20.0
	google.golang.org/grpc v1.43.0
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
)

require (
	github.com/envoyproxy/protoc-gen-validate v0.6.3 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.7.3 // indirect
	github.com/snovichkov/zap-gelf v1.0.1
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	golang.org/x/net v0.0.0-20210813160813-60bc85c4be6d // indirect
	golang.org/x/sys v0.0.0-20210816183151-1e6c022a8912 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20220118154757-00ab72f36ad5 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
)

replace github.com/tomazis/kioku/server/srv-frontend-api/pkg/srv-frontend-api => ./pkg/srv-frontend-api
