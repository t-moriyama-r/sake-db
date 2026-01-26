# バックエンド開発環境

## クイックスタート

### 初回セットアップ

```bash
cd backend
make setup
```

このコマンドで以下が自動的に実行されます：
- Lefthook（Git フック管理）のインストール
- golangci-lint（静的解析ツール）のインストール
- gqlgen（GraphQLコード生成）のインストール

### 開発フロー

通常通りコミットするだけで、自動的に静的解析が実行されます：

```bash
git add .
git commit -m "feat: 新機能を追加"
```

バックエンドファイルを変更した場合、`backend/.lefthook.yml`に定義された静的解析が自動実行されます。

#### GraphQLスキーマを変更した場合

スキーマファイル（`graph/schema/*.graphqls`）を変更した場合、フロントエンドの型定義を更新する必要があります。

コミット後に自動的に警告が表示されますが、以下のコマンドで簡単に更新できます：

```bash
# backend ディレクトリで実行
make update-frontend
```

このコマンドは：
1. バックエンドが起動しているかチェック
2. フロントエンドのスキーマを取得（`npm run fetch-schema`）
3. TypeScript型定義を生成（`npm run codegen`）

を自動的に実行します。

### 手動チェック

```bash
make check  # すべてのチェックを実行
make fmt    # コードフォーマット
make lint   # 静的解析
make vet    # Go vet
```

## 設定ファイル

- `.lefthook.yml`: バックエンド専用のLefthook設定
- `.golangci.yml`: golangci-lintの設定
- `Makefile`: 開発タスクの定義

## 詳細情報

詳しいセットアップ方法やトラブルシューティングは以下を参照：
- [バックエンド静的解析フックのセットアップ](../document/backend-hooks-setup.md)

