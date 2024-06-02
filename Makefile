#############
# go
#############
expor_path:
	@export PATH=$PATH:$GOPATH/bin
test:
	@go test -timeout 30s github.com/macabrabits/go_template/services -cover ./...
test2:
	@go test -timeout 30s -json > TestResults.json -cover ./...
vuln:
	# go install golang.org/x/vuln/cmd/govulncheck@latest
	@govulncheck ./...
sqlgen:
	# go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
	@sqlc generate
swag:
	# go install github.com/swaggo/swag/cmd/swag@latest
	@swag init
migrate:
	@docker run --rm -v $$PWD/db/schema:/migrations --network go_template_default migrate/migrate -path=/migrations/ -database 'mysql://root:root@tcp(mysql:3306)/app' up 1
rollback:
	@docker run --rm -v $$PWD/db/schema:/migrations --network go_template_default migrate/migrate -path=/migrations/ -database 'mysql://root:root@tcp(mysql:3306)/app' down 1

#############
# docker
#############	
up:
	@docker-compose up -d
up_debug:
	# go install github.com/go-delve/delve/cmd/dlv@latest
	@docker-compose -f docker-compose.yml -f docker-compose-debug.yml up -d
down:
	@docker-compose down



amend:
	@git commit -a --amend --no-edit && git push -f