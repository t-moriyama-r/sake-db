seeder:
	go run ./seeder.go

test:
	go test ./... -v

test-coverage:
	go test ./... -cover

test-coverage-html:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html

test-repo:
	go test ./db/repository/... -v

test-resolver:
	go test ./graph/resolver/... -v

benchmark:
	go test ./db/repository/liquorRepository -bench=. -benchmem