go mod init example/base
go run ./cmd/app
go build ./cmd/app

docker run -it --rm \
    -w "/app" \
    -e "air_wd=/app" \
    -v $(pwd)/cmd/app:/app \
    -p 8080:8080 \
    cosmtrek/air


docker-compose exec app sh -c "swag init"
docker-compose exec app bash

sudo chown dev:dev */*