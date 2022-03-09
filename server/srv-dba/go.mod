module github.com/tomazis/kioku/server/srv-dba

go 1.17

require (
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	github.com/snovichkov/zap-gelf v1.0.1
	github.com/tomazis/kioku/server/srv-dba/pkg/srv-dba v0.0.0-00010101000000-000000000000
	go.uber.org/zap v1.21.0
	google.golang.org/grpc v1.44.0
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
)

require github.com/pkg/errors v0.9.1 // indirect

require (
	github.com/envoyproxy/protoc-gen-validate v0.6.6 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.7.3 // indirect
	github.com/jmoiron/sqlx v1.3.4
	github.com/pressly/goose/v3 v3.5.3
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	golang.org/x/net v0.0.0-20220127200216-cd36cc0744dd // indirect
	golang.org/x/sys v0.0.0-20220209214540-3681064d5158 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20220302033224-9aa15565e42a // indirect
	google.golang.org/protobuf v1.27.1 // indirect
)

replace github.com/tomazis/kioku/server/srv-dba/pkg/srv-dba => ./pkg/srv-dba
