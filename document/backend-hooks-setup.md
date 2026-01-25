# バックエンド静的解析フックのセットアップ

このドキュメントでは、バックエンドコードの品質を保証するための静的解析フックの設定方法を説明します。

## 概要

このプロジェクトでは、コミット前に以下のチェックを自動実行します：

### フロントエンド
- **ESLint**: JavaScriptとVueファイルの構文チェック

### バックエンド
- **gofmt**: Goコードのフォーマットチェック
- **go vet**: Go標準の静的解析ツール
- **golangci-lint**: 高度なGo静的解析（複数のlinterを統合）
- **GraphQLスキーマ検証**: gqlgenによるスキーマの妥当性チェック

## 必要なツール

### 1. Lefthook
Git フック管理ツール

### 2. golangci-lint
Go言語の包括的な静的解析ツール

### 3. gqlgen
GraphQL スキーマからGoコードを生成するツール

## セットアップ手順

### 1. Lefthookのインストールとセットアップ

プロジェクトルートで以下を実行：

```bash
# 自動セットアップスクリプトを使用
./setup-hooks.sh
```

または手動でインストール：

```bash
# macOS
brew install lefthook

# Linux (Go経由)
go install github.com/evilmartians/lefthook@latest

# Lefthookのインストール（Git フックの有効化）
lefthook install
```

### 2. バックエンド開発ツールのインストール

バックエンドディレクトリで以下を実行：

```bash
cd backend
make install-tools
```

これにより以下がインストールされます：
- golangci-lint
- gqlgen

## 使い方

### 通常のコミットフロー

コミット時に自動的にチェックが実行されます：

```bash
git add .
git commit -m "feat: 新機能を追加"
# → 自動的に静的解析が実行されます
```

### フックをスキップする（非推奨）

緊急時のみ使用：

```bash
git commit -m "fix: 緊急修正" --no-verify
```

### 手動でチェックを実行

コミット前に手動でチェックすることもできます：

```bash
# バックエンドディレクトリで
cd backend

# フォーマットチェック
make fmt

# 静的解析
make lint

# Go vet
make vet

# GraphQL スキーマ検証
make generate

# すべてのチェック
make check
```

### 特定のフックのみ実行

```bash
# バックエンドのみ
lefthook run pre-commit --commands go-fmt,go-vet,go-lint

# フロントエンドのみ
lefthook run pre-commit --commands frontend-lint
```

## 静的解析ツールの詳細

### golangci-lint

設定ファイル: `backend/.golangci.yml`

有効なlinter：
- `govet`: Go標準のvetツール
- `errcheck`: エラーチェックの検証
- `staticcheck`: 静的解析
- `gosimple`: コードの簡略化提案
- `unused`: 未使用コードの検出
- `ineffassign`: 無効な代入の検出
- `typecheck`: 型チェック
- `gofmt`: コードフォーマット
- `goimports`: インポート整理

### GraphQLスキーマ検証

設定ファイル: `backend/gqlgen.yml`

- スキーマファイル: `backend/graph/schema/*.graphqls`
- gqlgenを使用してスキーマの妥当性を検証
- 構文エラーや型の不整合を検出

## トラブルシューティング

### golangci-lintが見つからない

```bash
cd backend
make install-tools
```

### Lefthookが動作しない

```bash
# 再インストール
lefthook install

# 設定の確認
lefthook run pre-commit --verbose
```

### GraphQL検証が失敗する

```bash
cd backend
# スキーマの再生成
make generate
```

### パフォーマンスが遅い

`lefthook.yml`で`parallel: true`が設定されているため、チェックは並列実行されます。
それでも遅い場合は、特定のlinterを無効化することを検討してください。

## ベストプラクティス

1. **コミット前に手動チェック**: 大きな変更の場合、コミット前に`make check`を実行
2. **小さなコミット**: 小さく頻繁にコミットすることで、チェック時間を短縮
3. **フック無視は避ける**: `--no-verify`の使用は最小限に
4. **定期的なツール更新**: 開発ツールを定期的に更新

## 参考リンク

- [Lefthook公式ドキュメント](https://github.com/evilmartians/lefthook)
- [golangci-lint公式ドキュメント](https://golangci-lint.run/)
- [gqlgen公式ドキュメント](https://gqlgen.com/)
