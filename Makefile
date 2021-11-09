watch:
	gow run ./cmd/server

run:
	go run ./cmd/server

serve:
	./bin/server

build:
	go build -v -x -o ./bin/server ./cmd/server

# Build including dynamic libraries
build-static:
	CGO_ENABLED=0 && GOOS=linux && GOARCH=amd64 && go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o ./bin/server_static ./cmd/server

build-docker:
	docker build -t jonamat/go-gin-gorm-rest-example:latest --no-cache .

push-docker:
	docker push jonamat/go-gin-gorm-rest-example:latest
