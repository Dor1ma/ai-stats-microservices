module github.com/Dor1ma/ai-stats-microservices/service2

go 1.23.6

require (
	github.com/Dor1ma/ai-stats-microservices/proto v0.0.0-00010101000000-000000000000
	github.com/joho/godotenv v1.5.1
	google.golang.org/grpc v1.70.0
)

require (
	golang.org/x/net v0.32.0 // indirect
	golang.org/x/sys v0.28.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20241202173237-19429a94021a // indirect
	google.golang.org/protobuf v1.36.5 // indirect
)

replace github.com/Dor1ma/ai-stats-microservices/proto => ../proto
