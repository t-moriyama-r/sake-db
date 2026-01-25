# バックエンドタスク

# シーダー実行
seeder:
	go run ./seeder.go

# Lefthookと開発ツールのセットアップ
setup:
	@echo "Setting up backend development environment..."
	@./setup-hooks.sh
	@$(MAKE) install-tools
	@echo "Setup completed! You can now commit with automatic static analysis."

# 依存関係のインストール
install-tools:
	@echo "Installing development tools..."
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install github.com/99designs/gqlgen@latest
	@echo "Tools installed successfully!"

# コードフォーマット（ファイルを変更）
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

# すべてのチェックを実行（注意：fmtはファイルを変更します）
check: vet lint generate
	@gofmt -l . | grep . && echo "Warning: Some files need formatting. Run 'make fmt' to fix." || echo "All checks passed!"

# 依存関係のダウンロード
deps:
	go mod download
	go mod tidy

.PHONY: seeder setup install-tools fmt lint vet generate check deps

