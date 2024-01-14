gotest:
	@go test ./...

gotestcov:
	@go test -short -coverpkg=./... -coverprofile=cover.out ./...
	@go tool cover -func cover.out

dbstatus:
	@goose -dir migrations postgres "host=postgres_16 password=root user=root dbname=postgres sslmode=disable" status

dbmigrate:
	@goose -dir migrations postgres "host=postgres_16 password=root user=root dbname=postgres sslmode=disable" up

dbrollback:
	@goose -dir migrations postgres "host=postgres_16 password=root user=root dbname=postgres sslmode=disable" down