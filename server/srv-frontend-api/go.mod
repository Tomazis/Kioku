module github.com/tomazis/kioku/server/srv-frontend-api

go 1.17

require (
	github.com/golang/mock v1.6.0
	github.com/tomazis/kioku/server/srv-frontend-api/pkg/srv-frontend-api v0.0.0-00010101000000-000000000000
	go.uber.org/zap v1.20.0
	google.golang.org/grpc v1.43.0
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
)

require (
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/go-task/slim-sprig v0.0.0-20210107165309-348f09dbbbc0 // indirect
	github.com/nxadm/tail v1.4.8 // indirect
	github.com/onsi/ginkgo v1.16.5 // indirect
	golang.org/x/tools v0.1.5 // indirect
	gopkg.in/tomb.v1 v1.0.0-20141024135613-dd632973f1e7 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

require (
	github.com/envoyproxy/protoc-gen-validate v0.6.3 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.7.3 // indirect
	github.com/onsi/ginkgo/v2 v2.0.0
	github.com/onsi/gomega v1.17.0
	github.com/snovichkov/zap-gelf v1.0.1
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	golang.org/x/net v0.0.0-20220114011407-0dd24b26b47d // indirect
	golang.org/x/sys v0.0.0-20220114195835-da31bd327af9 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20220118154757-00ab72f36ad5 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
)

replace github.com/tomazis/kioku/server/srv-frontend-api/pkg/srv-frontend-api => ./pkg/srv-frontend-api
