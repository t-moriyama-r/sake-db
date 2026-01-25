# バックエンド静的解析フックのセットアップ

このドキュメントでは、バックエンドコードの品質を保証するための静的解析フックの設定方法を説明します。

## 概要

バックエンドファイル専用の静的解析システムです。コミット前に以下のチェックを自動実行します：

### バックエンド
- **gofmt**: Goコードのフォーマットチェック
- **go vet**: Go標準の静的解析ツール
- **golangci-lint**: 高度なGo静的解析（複数のlinterを統合）
- **GraphQLスキーマ検証**: gqlgenによるスキーマの妥当性チェック

**重要**: Lefthookはリポジトリルートで動作しますが、`backend/`ディレクトリ内のファイルのみをチェックします。

## 必要なツール

### 1. Lefthook
Git フック管理ツール（リポジトリルートにインストール）

### 2. golangci-lint
Go言語の包括的な静的解析ツール

### 3. gqlgen
GraphQL スキーマからGoコードを生成するツール

## セットアップ手順

### クイックセットアップ（推奨）

バックエンドディレクトリで以下を実行：

```bash
cd backend
make setup
```

これだけで完了です！

### 手動セットアップ

#### 1. Lefthookのインストール

**macOSの場合:**
```bash
brew install lefthook
```

**Linuxまたはその他の環境:**
```bash
go install github.com/evilmartians/lefthook@latest
```

**Lefthookの有効化（リポジトリルートで実行）:**
```bash
# リポジトリルートに移動
cd /path/to/sake-db

# Lefthookをインストール
lefthook install
```

#### 2. バックエンド開発ツールのインストール

```bash
cd backend

# golangci-lintのインストール
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# gqlgenのインストール
go install github.com/99designs/gqlgen@latest
```

または、Makefileを使用：

```bash
cd backend
make install-tools
```

## 使い方

### 通常のコミットフロー

バックエンドディレクトリ内のファイルをコミットする際、自動的にチェックが実行されます：

```bash
# どこからでもコミット可能
git add backend/
git commit -m "feat: 新機能を追加"
# → 自動的にバックエンドファイルの静的解析が実行されます
```

チェックが失敗した場合、エラーメッセージが表示され、コミットはキャンセルされます。

### フックをスキップする（非推奨）

緊急時のみ使用：

```bash
git commit -m "fix: 緊急修正" --no-verify
```

⚠️ **注意**: これは推奨されません。コード品質の低下につながる可能性があります。

### 手動でチェックを実行

コミット前に手動でチェックすることもできます：

```bash
# バックエンドディレクトリで
cd backend

# フォーマットチェックと自動修正
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

### Lefthookの手動実行

```bash
# リポジトリルートで実行
lefthook run pre-commit

# または特定のコマンドのみ
lefthook run pre-commit --commands go-fmt
lefthook run pre-commit --commands go-vet
lefthook run pre-commit --commands go-lint
lefthook run pre-commit --commands graphql-validate
```

## 静的解析ツールの詳細

### golangci-lint

設定ファイル: `backend/.golangci.yml`

有効なlinter：
- `govet`: Go標準のvetツール - 一般的なバグを検出
- `errcheck`: エラーチェックの検証 - 未処理のエラーを検出
- `staticcheck`: 静的解析 - 高度な問題を検出
- `gosimple`: コードの簡略化提案 - より簡潔なコードを提案
- `unused`: 未使用コードの検出 - 使われていない変数や関数を検出
- `ineffassign`: 無効な代入の検出 - 効果のない代入を検出
- `typecheck`: 型チェック - 型の不整合を検出
- `gofmt`: コードフォーマット - コードスタイルをチェック
- `goimports`: インポート整理 - import文を整理

### GraphQLスキーマ検証

設定ファイル: `backend/gqlgen.yml`

- スキーマファイル: `backend/graph/schema/*.graphqls`
- gqlgenを使用してスキーマの妥当性を検証
- 構文エラーや型の不整合を検出
- 自動的にGraphQLリゾルバーコードを生成

## よくある問題と解決方法

### 問題: golangci-lintが見つからない

**解決策:**
```bash
cd backend
make install-tools
```

または手動で：
```bash
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
```

### 問題: Lefthookが動作しない

**解決策1: 再インストール**
```bash
# リポジトリルートで
lefthook uninstall
lefthook install
```

**解決策2: パスの確認**
```bash
# Go binディレクトリをPATHに追加
export PATH=$PATH:$(go env GOPATH)/bin

# .bashrc や .zshrc に追加することを推奨
echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.bashrc
```

**解決策3: 設定の確認**
```bash
# リポジトリルートで詳細モードで実行
lefthook run pre-commit --verbose
```

### 問題: GraphQL検証が失敗する

**解決策:**
```bash
cd backend

# スキーマの再生成
make generate

# または直接実行
go run github.com/99designs/gqlgen generate --config gqlgen.yml
```

### 問題: パフォーマンスが遅い

`lefthook.yml`で`parallel: true`が設定されているため、チェックは並列実行されます。

**さらに高速化する方法:**
1. 特定のlinterを無効化（`backend/.golangci.yml`を編集）
2. golangci-lintの`--new-from-rev`オプションを活用（既に設定済み）

### 問題: Goファイルのフォーマットエラー

**自動修正:**
```bash
cd backend
gofmt -w .
# または
make fmt
```

## ディレクトリ構成

```
/
├── lefthook.yml          # Lefthook設定（リポジトリルート）
└── backend/
    ├── .golangci.yml     # golangci-lint設定
    ├── setup-hooks.sh    # セットアップスクリプト
    ├── Makefile          # ビルドとチェックコマンド
    ├── gqlgen.yml        # GraphQL設定
    └── graph/
        └── schema/       # GraphQLスキーマファイル
```

## CI/CDとの統合

このフックはローカル開発環境用です。CI/CD環境では以下を実行してください：

```yaml
# GitHub Actions の例
- name: Run backend checks
  run: |
    cd backend
    make check
```

## ベストプラクティス

1. **初回セットアップ**: プロジェクトをクローンしたら`cd backend && make setup`を実行
2. **コミット前に手動チェック**: 大きな変更の場合、コミット前に`cd backend && make check`を実行
3. **小さなコミット**: 小さく頻繁にコミットすることで、チェック時間を短縮
4. **フック無視は避ける**: `--no-verify`の使用は最小限に
5. **定期的なツール更新**: 開発ツールを定期的に更新
   ```bash
   cd backend
   make install-tools
   ```
6. **エラーメッセージを読む**: linterのエラーメッセージには改善のヒントが含まれています
7. **チーム全体で使用**: 全員が同じツールを使用することでコード品質が統一されます

## Makefileコマンド一覧

バックエンドディレクトリで使用できるコマンド：

```bash
make setup         # Lefthookと開発ツールの完全セットアップ
make install-tools # 開発ツールのインストール
make fmt           # コードフォーマット
make lint          # golangci-lintの実行
make vet           # go vetの実行
make generate      # GraphQLコード生成
make check         # すべてのチェックを実行
make deps          # 依存関係の更新
```

## 参考リンク

- [Lefthook公式ドキュメント](https://github.com/evilmartians/lefthook)
- [golangci-lint公式ドキュメント](https://golangci-lint.run/)
- [gqlgen公式ドキュメント](https://gqlgen.com/)
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)

## サポート

問題が発生した場合は、以下を確認してください：

1. ツールが正しくインストールされているか
   ```bash
   which lefthook
   which golangci-lint
   ```

2. PATHに~/go/binが含まれているか
   ```bash
   echo $PATH | grep go/bin
   ```

3. リポジトリルートでlefthookが設定されているか
   ```bash
   cd /path/to/sake-db  # リポジトリルート
   lefthook run pre-commit --verbose
   ```



