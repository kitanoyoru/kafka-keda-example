gomod:
	go mod tidy && go mod vendor

lint:
	golangci-lint run -c .golangci.yaml

producer:
	go run ./cmd/main.go producer 

consumer:
	go run ./cmd/main.go consumer

gen-avro:
	avrogen -pkg user -o pkg/gen/go/user/user.go -tags json:snake,yaml:upper-camel ./schema/avro/user.avsc
