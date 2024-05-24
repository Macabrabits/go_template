go mod init example/base
go run ./cmd/app
go build ./cmd/app
go mod tidy
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
go test -timeout 30s github.com/macabrabits/go_template/services -cover ./...

docker run -v $PWD/db/schema:/migrations --network go_base_default migrate/migrate -path=/migrations/ -database 'mysql://root:root@tcp(mysql:3306)/app' up 2




