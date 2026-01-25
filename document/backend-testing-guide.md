# バックエンドテストガイド

## 概要

このドキュメントは、sake-dbプロジェクトのバックエンドテストに関する包括的なガイドです。
テストの構成、実行方法、技術選択の理由について説明します。

## 目次

1. [テスト戦略](#テスト戦略)
2. [技術選択](#技術選択)
3. [テスト構成](#テスト構成)
4. [テストの実行方法](#テストの実行方法)
5. [テストコードの書き方](#テストコードの書き方)
6. [ベストプラクティス](#ベストプラクティス)

## テスト戦略

### テストレベル

本プロジェクトでは、以下のテストレベルを採用しています：

1. **ユニットテスト（単体テスト）**
   - リポジトリ層の個別の関数をテスト
   - データベース操作の正確性を検証

2. **統合テスト（結合テスト）**
   - GraphQL Resolver層のテスト
   - リポジトリ層とResolver層の連携を検証

### テストスコープ

初期実装では、トップページの日本酒一覧取得機能に焦点を当てています：

- `randomRecommendList` GraphQLクエリ
- `GetRandomLiquors` リポジトリ関数
- 関連する補助関数

## 技術選択

### 1. Go標準のtestingパッケージ

**選択理由：**
- Goの標準ライブラリとして提供されており、追加の依存関係が不要
- シンプルで理解しやすいAPI
- Goコミュニティで広く採用されている標準

**使用方法：**
```go
func TestExampleFunction(t *testing.T) {
    // テストコード
}
```

### 2. testify/assert

**ライブラリ：** github.com/stretchr/testify

**選択理由：**
- 読みやすいアサーション（検証）を提供
- エラーメッセージが分かりやすい
- Goのテストコミュニティで広く使用されている

**使用例：**
```go
assert.Equal(t, expected, actual, "期待値と実際の値が一致すること")
assert.NotNil(t, result, "結果がnilでないこと")
```

### 3. testcontainers-go

**ライブラリ：** github.com/testcontainers/testcontainers-go

**選択理由：**
- 実際のMongoDBコンテナを使用した統合テストが可能
- テスト環境の完全な分離（他のテストやローカル環境に影響を与えない）
- テスト後の自動クリーンアップ
- CI/CD環境でも簡単に実行可能

**利点：**
- **実環境に近いテスト**: モックではなく実際のMongoDBを使用
- **再現性**: 各テストで新しいコンテナを起動するため、テスト間の依存関係がない
- **並列実行**: テストが独立しているため、並列実行が可能

## テスト構成

### ディレクトリ構造

```
backend/
├── db/
│   └── repository/
│       └── liquorRepository/
│           ├── repository.go          # リポジトリの実装
│           └── repository_test.go     # リポジトリのテスト
└── graph/
    └── resolver/
        ├── liquors.resolvers.go       # GraphQL Resolverの実装
        └── resolver_test.go           # Resolverのテスト
```

### テストファイルの命名規則

- テストファイル名: `<元のファイル名>_test.go`
- テスト関数名: `Test<関数名>_<テストケース>_<期待結果>`

**例：**
```go
// 正常系のテスト
func TestGetRandomLiquors_正常系_指定された件数が取得できること(t *testing.T) {}

// 異常系のテスト
func TestGetLiquorById_異常系_存在しないIDの場合エラーが返ること(t *testing.T) {}
```

## テストの実行方法

### 前提条件

- Docker/Podmanがインストールされていること
- テストコンテナを起動する権限があること

### すべてのテストを実行

```bash
# バックエンドディレクトリに移動
cd backend

# すべてのテストを実行
make test

# または直接goコマンドを使用
go test ./... -v
```

### 特定のテストを実行

```bash
# リポジトリ層のテストのみ
go test ./db/repository/liquorRepository -v

# Resolver層のテストのみ
go test ./graph/resolver -v

# 特定のテスト関数のみ
go test ./db/repository/liquorRepository -v -run TestGetRandomLiquors_正常系
```

### テストカバレッジの確認

```bash
# カバレッジを計測
go test ./... -cover

# 詳細なカバレッジレポートを生成
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out -o coverage.html

# ブラウザでカバレッジレポートを表示
open coverage.html  # macOS
xdg-open coverage.html  # Linux
```

### ベンチマークテストの実行

```bash
# ベンチマークテストを実行
go test ./db/repository/liquorRepository -bench=. -benchmem
```

## テストコードの書き方

### 基本的なテスト構造（AAA パターン）

テストは「Arrange（準備）」「Act（実行）」「Assert（検証）」の3つのフェーズで構成します。

```go
func TestGetRandomLiquors_正常系_指定された件数が取得できること(t *testing.T) {
    // Arrange（準備）: テスト用のデータベースをセットアップ
    testDB, cleanup := setupTestMongoDB(t)
    defer cleanup()
    
    repo := NewLiquorsRepository(testDB)
    insertTestLiquors(t, &repo, 10)
    
    // Act（実行）: テスト対象の関数を実行
    ctx := context.Background()
    result, err := repo.GetRandomLiquors(ctx, 5)
    
    // Assert（検証）: 結果を検証
    require.Nil(t, err, "エラーが発生してはいけません")
    assert.Equal(t, 5, len(result), "指定された件数の日本酒が取得できること")
}
```

### テストヘルパー関数

複雑なセットアップや共通処理は、ヘルパー関数として切り出します。

```go
// setupTestMongoDB はテスト用のMongoDBコンテナをセットアップする
func setupTestMongoDB(t *testing.T) (*db.DB, func()) {
    // MongoDBコンテナの起動
    // 接続設定
    // クリーンアップ関数の返却
}

// insertTestLiquors はテスト用の日本酒データを挿入する
func insertTestLiquors(t *testing.T, repo *LiquorsRepository, count int) []Model {
    // テストデータの生成と挿入
}
```

### エラーケースのテスト

エラーが期待される場合のテストも重要です。

```go
func TestGetLiquorById_異常系_存在しないIDの場合エラーが返ること(t *testing.T) {
    testDB, cleanup := setupTestMongoDB(t)
    defer cleanup()
    
    repo := NewLiquorsRepository(testDB)
    
    ctx := context.Background()
    nonExistentID := primitive.NewObjectID()
    result, err := repo.GetLiquorById(ctx, nonExistentID)
    
    // エラーが発生することを検証
    require.NotNil(t, err, "エラーが発生すること")
    assert.Nil(t, result, "結果がnilであること")
}
```

### テーブル駆動テスト（複数のケースをまとめてテスト）

同じロジックで複数のケースをテストする場合、テーブル駆動テストが効果的です。

```go
func TestGetRandomLiquors_境界値テスト(t *testing.T) {
    testDB, cleanup := setupTestMongoDB(t)
    defer cleanup()
    
    repo := NewLiquorsRepository(testDB)
    insertTestLiquors(t, &repo, 10)
    
    testCases := []struct {
        name     string
        limit    int
        expected int
    }{
        {"1件取得", 1, 1},
        {"5件取得", 5, 5},
        {"10件取得", 10, 10},
        {"全件より多く要求", 20, 10},
    }
    
    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            result, err := repo.GetRandomLiquors(context.Background(), tc.limit)
            require.Nil(t, err)
            assert.Equal(t, tc.expected, len(result), tc.name)
        })
    }
}
```

## ベストプラクティス

### 1. テストは独立させる

各テストは他のテストに依存せず、独立して実行できるようにします。

**良い例：**
```go
func TestA(t *testing.T) {
    testDB, cleanup := setupTestMongoDB(t)
    defer cleanup()
    // テストコード
}

func TestB(t *testing.T) {
    testDB, cleanup := setupTestMongoDB(t)
    defer cleanup()
    // テストコード
}
```

### 2. テスト名は日本語で具体的に

テスト名は、テストの内容が一目で分かるように具体的に書きます。

**良い例：**
```go
func TestGetRandomLiquors_正常系_指定された件数が取得できること(t *testing.T) {}
func TestGetRandomLiquors_異常系_データが0件の場合nilが返ること(t *testing.T) {}
```

**悪い例：**
```go
func TestGetRandomLiquors(t *testing.T) {}
func TestGetRandomLiquors2(t *testing.T) {}
```

### 3. 検証メッセージは日本語で明確に

アサーションのメッセージは、テストが失敗したときに原因が分かるように書きます。

**良い例：**
```go
assert.Equal(t, 5, len(result), "指定された件数の日本酒が取得できること")
assert.NotNil(t, result, "結果がnilであってはいけません")
```

### 4. クリーンアップは必ず実行

`defer` を使用して、テスト後のクリーンアップを確実に実行します。

```go
func TestExample(t *testing.T) {
    testDB, cleanup := setupTestMongoDB(t)
    defer cleanup()  // テスト終了時に必ず実行される
    
    // テストコード
}
```

### 5. テストデータは明示的に

テストで使用するデータは、テストコード内で明示的に定義します。

```go
func TestExample(t *testing.T) {
    // テストデータを明示的に定義
    testLiquor := liquorRepository.Model{
        ID:           primitive.NewObjectID(),
        Name:         "獺祭 純米大吟醸",
        CategoryID:   1,
        CategoryName: "純米大吟醸",
        // ...
    }
    
    // テストコード
}
```

### 6. require と assert の使い分け

- **require**: 失敗した場合、テストを即座に中断する
- **assert**: 失敗してもテストを続行する

```go
// requireは前提条件のチェックに使用
require.NoError(t, err, "データベース接続に失敗しました")
require.NotNil(t, result, "結果がnilであってはいけません")

// assertは複数の検証に使用
assert.Equal(t, expected, actual)
assert.True(t, condition)
```

## トラブルシューティング

### Docker/Podmanが起動できない

**問題：** testcontainersがコンテナを起動できない

**解決方法：**
1. Docker/Podmanが実行中か確認
   ```bash
   docker ps
   # または
   podman ps
   ```

2. 権限を確認
   ```bash
   docker info
   ```

3. testcontainersの環境変数を設定
   ```bash
   # Dockerの場合
   export DOCKER_HOST=unix:///var/run/docker.sock
   
   # Podmanの場合
   export DOCKER_HOST=unix:///run/user/$UID/podman/podman.sock
   export TESTCONTAINERS_RYUK_DISABLED=true
   ```

### テストが遅い

**問題：** テストの実行に時間がかかる

**解決方法：**
1. 並列実行を有効にする
   ```bash
   go test ./... -v -parallel 4
   ```

2. 特定のテストのみ実行
   ```bash
   go test ./... -v -short
   ```

### CI/CD環境でテストが失敗する

**問題：** ローカルでは成功するが、CI/CD環境で失敗する

**解決方法：**
1. CI環境でDockerが利用可能か確認
2. タイムアウト設定を調整
   ```bash
   go test ./... -v -timeout 10m
   ```

## テストカバレッジの目標

- **リポジトリ層**: 80%以上
- **Resolver層**: 70%以上
- **サービス層**: 75%以上（今後実装予定）

## 今後の拡張予定

1. **サービス層のテスト追加**
   - `liquorService` のテスト
   - ビジネスロジックのテスト

2. **認証・認可のテスト**
   - ユーザー認証のテスト
   - 権限チェックのテスト

3. **エンドツーエンドテスト**
   - GraphQLクエリの完全なテスト
   - フロントエンドとの統合テスト

4. **パフォーマンステスト**
   - ベンチマークテストの拡充
   - 負荷テストの追加

## 参考リンュース

- [Go Testing Package](https://pkg.go.dev/testing)
- [Testify Documentation](https://github.com/stretchr/testify)
- [Testcontainers for Go](https://golang.testcontainers.org/)
- [Go Test Best Practices](https://go.dev/doc/tutorial/add-a-test)

## まとめ

このテストガイドでは、sake-dbプロジェクトのバックエンドテストの基本的な方針と実装方法を説明しました。
テストは、コードの品質を保ち、リファクタリングを安全に行うための重要な要素です。

テストコードを書く際は、以下を心がけてください：

1. 読みやすく、理解しやすいテストコードを書く
2. テストケースは具体的で、明確な目的を持つ
3. 正常系だけでなく、異常系のテストも書く
4. テストは独立させ、再現性を保つ
5. 定期的にテストを実行し、品質を維持する

このガイドが、プロジェクトの品質向上に貢献することを願っています。
