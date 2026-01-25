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

### 手動チェック

```bash
make check  # すべてのチェックを実行
make fmt    # コードフォーマット
make lint   # 静的解析
make vet    # Go vet
```

## 詳細情報

詳しいセットアップ方法やトラブルシューティングは以下を参照：
- [バックエンド静的解析フックのセットアップ](../document/backend-hooks-setup.md)
