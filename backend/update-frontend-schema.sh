#!/bin/bash

# スクリプトのディレクトリを取得
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(dirname "$SCRIPT_DIR")"
FRONT_DIR="$PROJECT_ROOT/front"

# バックエンドのURL（環境変数で上書き可能）
BACKEND_URL="${BACKEND_URL:-http://localhost:8080}"

echo "================================================"
echo "フロントエンドのGraphQL型定義を更新します"
echo "================================================"
echo ""

# バックエンドが起動しているかチェック
echo "バックエンドの起動状態を確認中..."
echo "URL: $BACKEND_URL/query"
if ! curl -s "$BACKEND_URL/query" > /dev/null 2>&1; then
    echo ""
    echo "⚠️  エラー: バックエンドが起動していません！"
    echo ""
    echo "以下のいずれかの方法でバックエンドを起動してください："
    echo "  - Docker Compose: docker-compose up"
    echo "  - 直接実行: cd backend && go run main.go"
    echo ""
    echo "カスタムURLを使用する場合："
    echo "  BACKEND_URL=http://custom-url:8080 make update-frontend"
    echo ""
    exit 1
fi

echo "✓ バックエンドが起動しています"
echo ""

# フロントエンドディレクトリに移動
cd "$FRONT_DIR" || exit 1

# スキーマを取得
echo "1. GraphQLスキーマを取得中..."
if npm run fetch-schema; then
    echo "✓ スキーマの取得に成功しました"
else
    echo "✗ スキーマの取得に失敗しました"
    exit 1
fi

echo ""

# 型定義を生成
echo "2. TypeScript型定義を生成中..."
if npm run codegen; then
    echo "✓ 型定義の生成に成功しました"
else
    echo "✗ 型定義の生成に失敗しました"
    exit 1
fi

echo ""
echo "================================================"
echo "✓ フロントエンドの型定義の更新が完了しました！"
echo "================================================"
echo ""
echo "変更されたファイルを確認してください："
echo "  - front/schema.graphql"
echo "  - front/src/graphQL/auto-generated.ts"
echo ""
