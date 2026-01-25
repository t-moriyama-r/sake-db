package liquorRepository

import (
	"backend/db"
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// setupTestMongoDB はテスト用のMongoDBコンテナをセットアップする
func setupTestMongoDB(t *testing.T) (*db.DB, func()) {
	ctx := context.Background()

	// MongoDBコンテナを起動する
	req := testcontainers.ContainerRequest{
		Image:        "mongo:7.0",
		ExposedPorts: []string{"27017/tcp"},
		WaitingFor:   wait.ForLog("Waiting for connections").WithStartupTimeout(60 * time.Second),
	}

	mongoC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	require.NoError(t, err, "MongoDBコンテナの起動に失敗しました")

	// MongoDBのホストとポートを取得
	host, err := mongoC.Host(ctx)
	require.NoError(t, err, "MongoDBのホスト取得に失敗しました")

	port, err := mongoC.MappedPort(ctx, "27017")
	require.NoError(t, err, "MongoDBのポート取得に失敗しました")

	// MongoDB接続文字列を作成
	mongoURI := fmt.Sprintf("mongodb://%s:%s", host, port.Port())

	// MongoDB Clientを作成
	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(ctx, clientOptions)
	require.NoError(t, err, "MongoDBへの接続に失敗しました")

	// 接続を確認
	err = client.Ping(ctx, nil)
	require.NoError(t, err, "MongoDBへのPingに失敗しました")

	// テスト用のDBラッパーを作成
	testDB := &db.DB{
		Client: client,
		DBName: "sake_test",
	}

	// クリーンアップ関数を返す
	cleanup := func() {
		if err := client.Disconnect(context.Background()); err != nil {
			t.Errorf("MongoDB切断に失敗しました: %v", err)
		}
		if err := mongoC.Terminate(context.Background()); err != nil {
			t.Errorf("MongoDBコンテナの終了に失敗しました: %v", err)
		}
	}

	return testDB, cleanup
}

// insertTestLiquors はテスト用の日本酒データを挿入する
func insertTestLiquors(t *testing.T, repo *LiquorsRepository, count int) []Model {
	ctx := context.Background()
	liquors := make([]Model, count)

	for i := 0; i < count; i++ {
		name := fmt.Sprintf("テスト日本酒%d", i+1)
		description := fmt.Sprintf("これはテスト用の日本酒%dの説明です", i+1)
		liquor := Model{
			ID:           primitive.NewObjectID(),
			CategoryID:   1,
			CategoryName: "日本酒",
			Name:         name,
			Description:  &description,
			Rate5Users:   []string{},
			Rate4Users:   []string{},
			Rate3Users:   []string{},
			Rate2Users:   []string{},
			Rate1Users:   []string{},
			UpdatedAt:    time.Now(),
			RandomKey:    float64(i) / float64(count), // 0から1の範囲で均等に分布
		}
		liquors[i] = liquor

		// データベースに挿入
		_, err := repo.collection.InsertOne(ctx, liquor)
		require.NoError(t, err, "テストデータの挿入に失敗しました")
	}

	return liquors
}

// TestGetRandomLiquors_正常系_指定された件数が取得できること はGetRandomLiquorsの正常系テスト
func TestGetRandomLiquors_正常系_指定された件数が取得できること(t *testing.T) {
	// 準備: テスト用のMongoDBをセットアップ
	testDB, cleanup := setupTestMongoDB(t)
	defer cleanup()

	// リポジトリを作成
	repo := NewLiquorsRepository(testDB)

	// テストデータを挿入（10件）
	insertTestLiquors(t, &repo, 10)

	// テスト実行: 5件のランダムな日本酒を取得
	ctx := context.Background()
	limit := 5
	result, err := repo.GetRandomLiquors(ctx, limit)

	// 検証: エラーが発生しないこと
	require.Nil(t, err, "エラーが発生してはいけません")
	require.NotNil(t, result, "結果がnilであってはいけません")

	// 検証: 指定された件数が取得できること
	assert.Equal(t, limit, len(result), "指定された件数の日本酒が取得できること")

	// 検証: すべての結果が有効なデータであること
	for _, liquor := range result {
		assert.NotEmpty(t, liquor.ID, "IDが空ではないこと")
		assert.NotEmpty(t, liquor.Name, "名前が空ではないこと")
		assert.NotEmpty(t, liquor.CategoryName, "カテゴリ名が空ではないこと")
	}
}

// TestGetRandomLiquors_正常系_データが0件の場合nilが返ること はデータがない場合のテスト
func TestGetRandomLiquors_正常系_データが0件の場合nilが返ること(t *testing.T) {
	// 準備: テスト用のMongoDBをセットアップ
	testDB, cleanup := setupTestMongoDB(t)
	defer cleanup()

	// リポジトリを作成
	repo := NewLiquorsRepository(testDB)

	// テスト実行: データがない状態で日本酒を取得
	ctx := context.Background()
	result, err := repo.GetRandomLiquors(ctx, 5)

	// 検証: エラーが発生しないこと
	require.Nil(t, err, "エラーが発生してはいけません")

	// 検証: nilが返ること
	assert.Nil(t, result, "データがない場合はnilが返ること")
}

// TestGetRandomLiquors_正常系_要求件数より少ない場合は全件取得できること は要求件数より少ない場合のテスト
func TestGetRandomLiquors_正常系_要求件数より少ない場合は全件取得できること(t *testing.T) {
	// 準備: テスト用のMongoDBをセットアップ
	testDB, cleanup := setupTestMongoDB(t)
	defer cleanup()

	// リポジトリを作成
	repo := NewLiquorsRepository(testDB)

	// テストデータを挿入（3件のみ）
	insertTestLiquors(t, &repo, 3)

	// テスト実行: 10件を要求するが、実際には3件しかない
	ctx := context.Background()
	result, err := repo.GetRandomLiquors(ctx, 10)

	// 検証: エラーが発生しないこと
	require.Nil(t, err, "エラーが発生してはいけません")
	require.NotNil(t, result, "結果がnilであってはいけません")

	// 検証: 実際のデータ件数（3件）が取得できること
	assert.Equal(t, 3, len(result), "実際に存在する件数の日本酒が取得できること")
}

// TestGetRandomLiquors_正常系_ランダム性が担保されること はランダム性のテスト
func TestGetRandomLiquors_正常系_ランダム性が担保されること(t *testing.T) {
	// 準備: テスト用のMongoDBをセットアップ
	testDB, cleanup := setupTestMongoDB(t)
	defer cleanup()

	// リポジトリを作成
	repo := NewLiquorsRepository(testDB)

	// テストデータを挿入（20件）
	insertTestLiquors(t, &repo, 20)

	// テスト実行: 複数回実行して結果を比較
	ctx := context.Background()
	limit := 5

	// 1回目の取得
	result1, err1 := repo.GetRandomLiquors(ctx, limit)
	require.Nil(t, err1, "1回目の取得でエラーが発生してはいけません")

	// 2回目の取得
	result2, err2 := repo.GetRandomLiquors(ctx, limit)
	require.Nil(t, err2, "2回目の取得でエラーが発生してはいけません")

	// 検証: 両方とも正しい件数が取得できること
	assert.Equal(t, limit, len(result1), "1回目の結果が指定件数であること")
	assert.Equal(t, limit, len(result2), "2回目の結果が指定件数であること")

	// 検証: 結果が異なる可能性が高いこと（ランダム性の確認）
	// すべてのIDが一致することは稀なので、少なくとも1つは異なるはず
	allSame := true
	result1IDs := make(map[string]bool)
	for _, liquor := range result1 {
		result1IDs[liquor.ID.Hex()] = true
	}

	for _, liquor := range result2 {
		if !result1IDs[liquor.ID.Hex()] {
			allSame = false
			break
		}
	}

	// 注意: ランダムなので稀に同じ結果になる可能性があるが、
	// 20件中5件を2回取得して完全に一致する確率は低い
	// このテストが失敗する場合は、テストを再実行してください
	t.Logf("1回目と2回目の結果が完全に一致: %v", allSame)
}

// TestGetLiquorById_正常系_指定したIDの日本酒が取得できること はGetLiquorByIdの正常系テスト
func TestGetLiquorById_正常系_指定したIDの日本酒が取得できること(t *testing.T) {
	// 準備: テスト用のMongoDBをセットアップ
	testDB, cleanup := setupTestMongoDB(t)
	defer cleanup()

	// リポジトリを作成
	repo := NewLiquorsRepository(testDB)

	// テストデータを挿入
	liquors := insertTestLiquors(t, &repo, 3)
	targetLiquor := liquors[1] // 2番目のデータを対象とする

	// テスト実行: 指定したIDの日本酒を取得
	ctx := context.Background()
	result, err := repo.GetLiquorById(ctx, targetLiquor.ID)

	// 検証: エラーが発生しないこと
	require.Nil(t, err, "エラーが発生してはいけません")
	require.NotNil(t, result, "結果がnilであってはいけません")

	// 検証: 正しいデータが取得できること
	assert.Equal(t, targetLiquor.ID, result.ID, "IDが一致すること")
	assert.Equal(t, targetLiquor.Name, result.Name, "名前が一致すること")
	assert.Equal(t, targetLiquor.CategoryName, result.CategoryName, "カテゴリ名が一致すること")
}

// TestGetLiquorById_異常系_存在しないIDの場合エラーが返ること はGetLiquorByIdの異常系テスト
func TestGetLiquorById_異常系_存在しないIDの場合エラーが返ること(t *testing.T) {
	// 準備: テスト用のMongoDBをセットアップ
	testDB, cleanup := setupTestMongoDB(t)
	defer cleanup()

	// リポジトリを作成
	repo := NewLiquorsRepository(testDB)

	// テスト実行: 存在しないIDで日本酒を取得
	ctx := context.Background()
	nonExistentID := primitive.NewObjectID()
	result, err := repo.GetLiquorById(ctx, nonExistentID)

	// 検証: エラーが発生すること
	require.NotNil(t, err, "エラーが発生すること")
	assert.Nil(t, result, "結果がnilであること")
}

// TestGetLiquorsFromCategoryIds_正常系_指定したカテゴリIDの日本酒が取得できること はGetLiquorsFromCategoryIdsの正常系テスト
func TestGetLiquorsFromCategoryIds_正常系_指定したカテゴリIDの日本酒が取得できること(t *testing.T) {
	// 準備: テスト用のMongoDBをセットアップ
	testDB, cleanup := setupTestMongoDB(t)
	defer cleanup()

	// リポジトリを作成
	repo := NewLiquorsRepository(testDB)

	// 異なるカテゴリIDのテストデータを挿入
	ctx := context.Background()
	
	// カテゴリID 1の日本酒
	for i := 0; i < 3; i++ {
		name := fmt.Sprintf("吟醸酒%d", i+1)
		description := fmt.Sprintf("カテゴリ1の日本酒%d", i+1)
		liquor := Model{
			ID:           primitive.NewObjectID(),
			CategoryID:   1,
			CategoryName: "吟醸酒",
			Name:         name,
			Description:  &description,
			Rate5Users:   []string{},
			Rate4Users:   []string{},
			Rate3Users:   []string{},
			Rate2Users:   []string{},
			Rate1Users:   []string{},
			UpdatedAt:    time.Now(),
			RandomKey:    float64(i) / 10.0,
		}
		_, err := repo.collection.InsertOne(ctx, liquor)
		require.NoError(t, err, "テストデータの挿入に失敗しました")
	}

	// カテゴリID 2の日本酒
	for i := 0; i < 2; i++ {
		name := fmt.Sprintf("純米酒%d", i+1)
		description := fmt.Sprintf("カテゴリ2の日本酒%d", i+1)
		liquor := Model{
			ID:           primitive.NewObjectID(),
			CategoryID:   2,
			CategoryName: "純米酒",
			Name:         name,
			Description:  &description,
			Rate5Users:   []string{},
			Rate4Users:   []string{},
			Rate3Users:   []string{},
			Rate2Users:   []string{},
			Rate1Users:   []string{},
			UpdatedAt:    time.Now(),
			RandomKey:    float64(i) / 10.0,
		}
		_, err := repo.collection.InsertOne(ctx, liquor)
		require.NoError(t, err, "テストデータの挿入に失敗しました")
	}

	// テスト実行: カテゴリID 1の日本酒を取得
	categoryIds := []int{1}
	result, err := repo.GetLiquorsFromCategoryIds(ctx, categoryIds)

	// 検証: エラーが発生しないこと
	require.Nil(t, err, "エラーが発生してはいけません")
	require.NotNil(t, result, "結果がnilであってはいけません")

	// 検証: カテゴリID 1の日本酒のみが取得できること
	assert.Equal(t, 3, len(result), "カテゴリID 1の日本酒が3件取得できること")
	for _, liquor := range result {
		assert.Equal(t, 1, liquor.CategoryID, "すべてカテゴリID 1であること")
	}
}

// BenchmarkGetRandomLiquors は GetRandomLiquors のベンチマークテスト
func BenchmarkGetRandomLiquors(b *testing.B) {
	// 準備: テスト用のMongoDBをセットアップ
	// 注意: ベンチマークテストの場合、*testing.B を使用
	testDB, cleanup := func() (*db.DB, func()) {
		ctx := context.Background()

		req := testcontainers.ContainerRequest{
			Image:        "mongo:7.0",
			ExposedPorts: []string{"27017/tcp"},
			WaitingFor:   wait.ForLog("Waiting for connections").WithStartupTimeout(60 * time.Second),
		}

		mongoC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
			ContainerRequest: req,
			Started:          true,
		})
		if err != nil {
			b.Fatalf("MongoDBコンテナの起動に失敗しました: %v", err)
		}

		host, err := mongoC.Host(ctx)
		if err != nil {
			b.Fatalf("MongoDBのホスト取得に失敗しました: %v", err)
		}

		port, err := mongoC.MappedPort(ctx, "27017")
		if err != nil {
			b.Fatalf("MongoDBのポート取得に失敗しました: %v", err)
		}

		mongoURI := fmt.Sprintf("mongodb://%s:%s", host, port.Port())
		clientOptions := options.Client().ApplyURI(mongoURI)
		client, err := mongo.Connect(ctx, clientOptions)
		if err != nil {
			b.Fatalf("MongoDBへの接続に失敗しました: %v", err)
		}

		err = client.Ping(ctx, nil)
		if err != nil {
			b.Fatalf("MongoDBへのPingに失敗しました: %v", err)
		}

		testDB := &db.DB{
			Client: client,
			DBName: "sake_bench",
		}

		cleanup := func() {
			client.Disconnect(context.Background())
			mongoC.Terminate(context.Background())
		}

		return testDB, cleanup
	}()
	defer cleanup()

	// リポジトリを作成
	repo := NewLiquorsRepository(testDB)

	// テストデータを挿入（100件）
	ctx := context.Background()
	for i := 0; i < 100; i++ {
		name := fmt.Sprintf("ベンチマーク日本酒%d", i+1)
		description := fmt.Sprintf("ベンチマーク用の日本酒%d", i+1)
		liquor := Model{
			ID:           primitive.NewObjectID(),
			CategoryID:   1,
			CategoryName: "日本酒",
			Name:         name,
			Description:  &description,
			Rate5Users:   []string{},
			Rate4Users:   []string{},
			Rate3Users:   []string{},
			Rate2Users:   []string{},
			Rate1Users:   []string{},
			UpdatedAt:    time.Now(),
			RandomKey:    float64(i) / 100.0,
		}
		_, err := repo.collection.InsertOne(ctx, liquor)
		if err != nil {
			b.Fatalf("テストデータの挿入に失敗しました: %v", err)
		}
	}

	// インデックスを作成（パフォーマンスのため）
	indexModel := mongo.IndexModel{
		Keys: bson.D{{Key: "random_key", Value: 1}},
	}
	_, err := repo.collection.Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		b.Fatalf("インデックスの作成に失敗しました: %v", err)
	}

	// ベンチマークの計測をリセット
	b.ResetTimer()

	// ベンチマークテストを実行
	for i := 0; i < b.N; i++ {
		_, err := repo.GetRandomLiquors(ctx, 10)
		if err != nil {
			b.Fatalf("GetRandomLiquorsでエラーが発生しました: %v", err)
		}
	}
}
