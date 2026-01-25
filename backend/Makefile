# バックエンドタスク

# シーダー実行
seeder:
	go run ./seeder.go

# 依存関係のインストール
install-tools:
	@echo "Installing development tools..."
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install github.com/99designs/gqlgen@latest
	@echo "Tools installed successfully!"

# コードフォーマット
fmt:
	gofmt -w .

# Lintチェック
lint:
	golangci-lint run ./...

# Go vet
vet:
	go vet ./...

# GraphQLスキーマ検証と生成
generate:
	go run github.com/99designs/gqlgen generate --config gqlgen.yml

# すべてのチェックを実行
check: fmt vet lint generate
	@echo "All checks passed!"

# 依存関係のダウンロード
deps:
	go mod download
	go mod tidy

.PHONY: seeder install-tools fmt lint vet generate check deps
