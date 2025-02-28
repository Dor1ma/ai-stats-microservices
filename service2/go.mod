module github.com/Dor1ma/ai-stats-microservices/service2

go 1.23.6

require (
	github.com/Dor1ma/ai-stats-microservices/proto v0.0.0-00010101000000-000000000000
	github.com/joho/godotenv v1.5.1
	github.com/stretchr/testify v1.8.1
	go.uber.org/zap v1.27.0
	google.golang.org/grpc v1.70.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/stretchr/objx v0.5.0 // indirect
	go.uber.org/multierr v1.10.0 // indirect
	golang.org/x/net v0.32.0 // indirect
	golang.org/x/sys v0.28.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20241202173237-19429a94021a // indirect
	google.golang.org/protobuf v1.36.5 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/Dor1ma/ai-stats-microservices/proto => ../proto
