# GraphQLスキーマ変更時の自動化ガイド

## 概要

バックエンドのGraphQLスキーマ（`backend/graph/schema/*.graphqls`）を変更した場合、フロントエンドの型定義を更新する必要があります。

この自動化により、スキーマ変更時に必要な手順が明確化され、開発者が忘れずに実行できるようになりました。

## 実装内容

### 1. Copilot/AI向けのルール（`.github/copilot-instructions.md`）

GitHub CopilotやClaude CodeなどのAIツールが、バックエンドスキーマを変更する際に必ず実行すべき手順を明記しました。

- スキーマ変更時の必須手順を明確化
- AIツール使用時も同じルールを適用するよう指示

### 2. Git フック（`backend/.lefthook.yml`）

スキーマファイルをコミットした際に、自動的に警告メッセージを表示します。

**動作:**
- `backend/graph/schema/*.graphqls`ファイルをコミットした際に自動検出
- フロントエンド型定義の更新手順を表示
- 実行を忘れた場合のリスクを警告

### 3. ヘルパースクリプト（`backend/update-frontend-schema.sh`）

フロントエンドの型定義を簡単に更新できる自動化スクリプトを作成しました。

**機能:**
- バックエンドが起動しているかチェック
- `npm run fetch-schema`を自動実行
- `npm run codegen`を自動実行
- 各ステップの成功/失敗を表示

### 4. Makeコマンド（`backend/Makefile`）

簡単に実行できるMakeコマンドを追加しました。

```bash
make update-frontend
```

### 5. ドキュメント更新

- `backend/README.md`: クイックスタートガイドに追加
- `document/backend.md`: 開発手順に詳細を追加

## 使い方

### スキーマ変更後の手順

バックエンドのスキーマファイルを変更した後、以下のいずれかの方法で実行します。

#### 方法1: 自動スクリプトを使用（推奨）

```bash
cd backend
make update-frontend
```

このコマンドは以下を自動で実行します：
1. バックエンドが起動しているかチェック
2. フロントエンドのスキーマを取得
3. TypeScript型定義を生成

#### 方法2: 手動で実行

```bash
# 1. バックエンドを起動（別ターミナルで）
docker-compose up
# または
cd backend && go run main.go

# 2. フロントエンドディレクトリに移動
cd front

# 3. スキーマを取得
npm run fetch-schema

# 4. 型定義を生成
npm run codegen
```

### Git コミット時の自動警告

スキーマファイルをコミットすると、自動的に以下のような警告が表示されます：

```
⚠️  ==========================================
⚠️  GraphQLスキーマが変更されました！
⚠️  ==========================================

フロントエンドの型定義を更新してください：

  1. バックエンドを起動してください（まだの場合）
  2. cd front
  3. npm run fetch-schema
  4. npm run codegen

これらのコマンドを実行しないと、フロントエンドで
ビルドエラーや型エラーが発生する可能性があります。
```

## トラブルシューティング

### エラー: バックエンドが起動していません

**原因:**
- バックエンドサーバーが起動していない
- ポート8080でリッスンしていない
- カスタムポートを使用している

**解決方法:**
```bash
# Docker Composeで起動
docker-compose up

# または直接実行
cd backend && go run main.go

# カスタムURLを使用する場合
BACKEND_URL=http://custom-url:9000 make update-frontend
```

### エラー: スキーマの取得に失敗しました

**原因:**
- バックエンドのGraphQLエンドポイントにアクセスできない
- ネットワーク接続の問題

**解決方法:**
1. バックエンドが起動していることを確認
2. `http://localhost:8080/query`にアクセスできることを確認
3. ファイアウォール設定を確認

### エラー: 型定義の生成に失敗しました

**原因:**
- GraphQLスキーマに構文エラーがある
- codegenの設定に問題がある

**解決方法:**
1. `backend/graph/schema/*.graphqls`の構文を確認
2. `front/codegen.ts`の設定を確認
3. `npm install`で依存関係を再インストール

## 関連ファイル

- `.github/copilot-instructions.md`: AI向けのルール定義
- `backend/.lefthook.yml`: Git フック設定
- `backend/update-frontend-schema.sh`: 自動化スクリプト
- `backend/Makefile`: Makeコマンド定義
- `backend/README.md`: バックエンド開発ガイド
- `document/backend.md`: 詳細な開発手順
- `front/codegen.ts`: GraphQL Code Generator設定
- `front/package.json`: npm スクリプト定義

## 今後の改善案

1. CI/CDパイプラインへの統合
   - PRマージ時に自動的にスキーマ同期を確認

2. pre-commit フックの追加
   - コミット前にスキーマの整合性をチェック

3. VSCode拡張機能の検討
   - スキーマ変更を検出して自動的に型定義を更新
