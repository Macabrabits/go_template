go mod init example/base
go run ./cmd/app
go build ./cmd/app
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
go mod tidy github.com/sqlc-dev/sqlc/cmd/sqlc@latest
export PATH=$PATH:$GOPATH/bin


docker run -it --rm \
    -w "/app" \
    -e "air_wd=/app" \
    -v $(pwd)/cmd/app:/app \
    -p 8080:8080 \
    cosmtrek/air


docker-compose exec app sh -c "swag init"
docker-compose exec app bash

sudo chown dev:dev */*



docker-compose -f docker-compose.yml -f docker-compose-debug.yml up -d

go test -timeout 30s github.com/macabrabits/go_template/services -cover
