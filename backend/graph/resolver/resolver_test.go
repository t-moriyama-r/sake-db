package resolver

import (
	"backend/db"
	"backend/db/repository/bookmarkRepository"
	"backend/db/repository/categoriesRepository"
	"backend/db/repository/flavorMapRepository"
	"backend/db/repository/liquorRepository"
	"backend/db/repository/userRepository"
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// setupTestDatabase はテスト用のデータベース環境をセットアップする
func setupTestDatabase(t *testing.T) (*db.DB, func()) {
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
		DBName: "sake_resolver_test",
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

// setupTestResolver はテスト用のResolverをセットアップする
func setupTestResolver(t *testing.T, testDB *db.DB) *Resolver {
	// 各リポジトリを初期化
	liquorRepo := liquorRepository.NewLiquorsRepository(testDB)
	categoryRepo := categoriesRepository.NewCategoryRepository(testDB)
	userRepo := userRepository.NewUsersRepository(testDB)
	bookmarkRepo := bookmarkRepository.NewBookMarkRepository(testDB)
	flavorMapMstRepo := flavorMapRepository.NewFlavorMapMasterRepository(testDB)
	flavorMapRepo := flavorMapRepository.NewFlavorMapRepository(testDB)
	flavorLiqRepo := flavorMapRepository.NewFlavorToLiquorRepository(testDB)

	// MongoDBのDatabaseインスタンスを取得
	database := testDB.Client.Database(testDB.DBName)

	// Resolverを作成
	resolver := &Resolver{
		DB:               database,
		LiquorRepo:       liquorRepo,
		CategoryRepo:     categoryRepo,
		UserRepo:         userRepo,
		BookmarkRepo:     bookmarkRepo,
		FlavorMapRepo:    flavorMapRepo,
		FlavorMapMstRepo: flavorMapMstRepo,
		FlavorLiqRepo:    flavorLiqRepo,
	}

	return resolver
}

// insertTestLiquorsForResolver はテスト用の日本酒データを挿入する（Resolver用）
func insertTestLiquorsForResolver(t *testing.T, repo liquorRepository.LiquorsRepository, count int) []liquorRepository.Model {
	ctx := context.Background()
	liquors := make([]liquorRepository.Model, count)

	for i := 0; i < count; i++ {
		name := fmt.Sprintf("リゾルバーテスト日本酒%d", i+1)
		description := fmt.Sprintf("これはリゾルバーテスト用の日本酒%dの説明です。芳醇な香りと深い味わいが特徴です。", i+1)
		versionNo := 1 // VersionNoを設定
		liquor := liquorRepository.Model{
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
			RandomKey:    float64(i) / float64(count),
			VersionNo:    &versionNo,
		}
		liquors[i] = liquor

		// データベースに挿入
		collection := repo.DB.Collection(liquorRepository.CollectionName)
		_, err := collection.InsertOne(ctx, liquor)
		require.NoError(t, err, "テストデータの挿入に失敗しました")
	}

	return liquors
}

// TestRandomRecommendList_正常系_指定された件数の日本酒リストが取得できること はRandomRecommendListの正常系テスト
func TestRandomRecommendList_正常系_指定された件数の日本酒リストが取得できること(t *testing.T) {
	// 準備: テスト用のデータベースをセットアップ
	testDB, cleanup := setupTestDatabase(t)
	defer cleanup()

	// 準備: Resolverをセットアップ
	resolver := setupTestResolver(t, testDB)

	// 準備: テストデータを挿入（15件）
	insertTestLiquorsForResolver(t, resolver.LiquorRepo, 15)

	// テスト実行: 10件のランダムな日本酒リストを取得
	ctx := context.Background()
	limit := 10
	queryResolver := resolver.Query()
	result, err := queryResolver.RandomRecommendList(ctx, limit)

	// 検証: エラーが発生しないこと
	require.Nil(t, err, "エラーが発生してはいけません")
	require.NotNil(t, result, "結果がnilであってはいけません")

	// 検証: 指定された件数の日本酒が取得できること
	assert.Equal(t, limit, len(result), "指定された件数の日本酒が取得できること")

	// 検証: 各日本酒のデータが正しく変換されていること
	for _, liquor := range result {
		// IDが正しくHex文字列に変換されていること
		assert.NotEmpty(t, liquor.ID, "IDが空ではないこと")
		assert.Len(t, liquor.ID, 24, "IDが24文字のHex文字列であること")

		// 必須フィールドが設定されていること
		assert.NotEmpty(t, liquor.Name, "名前が空ではないこと")
		assert.NotEmpty(t, liquor.CategoryName, "カテゴリ名が空ではないこと")
		assert.Equal(t, 1, liquor.CategoryID, "カテゴリIDが正しいこと")

		// オプショナルフィールドが設定されていること
		assert.NotNil(t, liquor.Description, "説明が設定されていること")

		// 配列フィールドが初期化されていること
		assert.NotNil(t, liquor.Rate5Users, "評価5のユーザーリストが初期化されていること")
		assert.NotNil(t, liquor.Rate4Users, "評価4のユーザーリストが初期化されていること")
		assert.NotNil(t, liquor.Rate3Users, "評価3のユーザーリストが初期化されていること")
		assert.NotNil(t, liquor.Rate2Users, "評価2のユーザーリストが初期化されていること")
		assert.NotNil(t, liquor.Rate1Users, "評価1のユーザーリストが初期化されていること")

		// 更新日時が設定されていること
		assert.NotEmpty(t, liquor.UpdatedAt, "更新日時が設定されていること")
	}
}

// TestRandomRecommendList_正常系_データが0件の場合nilまたは空配列が返ること はデータがない場合のテスト
func TestRandomRecommendList_正常系_データが0件の場合nilまたは空配列が返ること(t *testing.T) {
	// 準備: テスト用のデータベースをセットアップ
	testDB, cleanup := setupTestDatabase(t)
	defer cleanup()

	// 準備: Resolverをセットアップ（テストデータは挿入しない）
	resolver := setupTestResolver(t, testDB)

	// テスト実行: データがない状態で日本酒リストを取得
	ctx := context.Background()
	queryResolver := resolver.Query()
	result, err := queryResolver.RandomRecommendList(ctx, 10)

	// 検証: エラーが発生しないこと
	require.Nil(t, err, "エラーが発生してはいけません")

	// 検証: nilまたは空配列が返ること
	if result != nil {
		assert.Empty(t, result, "データがない場合は空配列が返ること")
	}
	// nilの場合もOK
}

// TestRandomRecommendList_正常系_要求件数より少ない場合は全件取得できること は要求件数より少ない場合のテスト
func TestRandomRecommendList_正常系_要求件数より少ない場合は全件取得できること(t *testing.T) {
	// 準備: テスト用のデータベースをセットアップ
	testDB, cleanup := setupTestDatabase(t)
	defer cleanup()

	// 準備: Resolverをセットアップ
	resolver := setupTestResolver(t, testDB)

	// 準備: テストデータを挿入（5件のみ）
	insertTestLiquorsForResolver(t, resolver.LiquorRepo, 5)

	// テスト実行: 20件を要求するが、実際には5件しかない
	ctx := context.Background()
	queryResolver := resolver.Query()
	result, err := queryResolver.RandomRecommendList(ctx, 20)

	// 検証: エラーが発生しないこと
	require.Nil(t, err, "エラーが発生してはいけません")
	require.NotNil(t, result, "結果がnilであってはいけません")

	// 検証: 実際のデータ件数（5件）が取得できること
	assert.Equal(t, 5, len(result), "実際に存在する件数の日本酒が取得できること")
}

// TestRandomRecommendList_正常系_limit1の場合でも正常に動作すること は最小のlimitでのテスト
func TestRandomRecommendList_正常系_limit1の場合でも正常に動作すること(t *testing.T) {
	// 準備: テスト用のデータベースをセットアップ
	testDB, cleanup := setupTestDatabase(t)
	defer cleanup()

	// 準備: Resolverをセットアップ
	resolver := setupTestResolver(t, testDB)

	// 準備: テストデータを挿入
	insertTestLiquorsForResolver(t, resolver.LiquorRepo, 10)

	// テスト実行: 1件のみ取得
	ctx := context.Background()
	queryResolver := resolver.Query()
	result, err := queryResolver.RandomRecommendList(ctx, 1)

	// 検証: エラーが発生しないこと
	require.Nil(t, err, "エラーが発生してはいけません")
	require.NotNil(t, result, "結果がnilであってはいけません")

	// 検証: 1件のみ取得できること
	assert.Equal(t, 1, len(result), "1件の日本酒が取得できること")

	// 検証: データの内容が正しいこと
	liquor := result[0]
	assert.NotEmpty(t, liquor.ID, "IDが設定されていること")
	assert.NotEmpty(t, liquor.Name, "名前が設定されていること")
}

// TestRandomRecommendList_正常系_複数回実行しても正常に動作すること は繰り返し実行のテスト
func TestRandomRecommendList_正常系_複数回実行しても正常に動作すること(t *testing.T) {
	// 準備: テスト用のデータベースをセットアップ
	testDB, cleanup := setupTestDatabase(t)
	defer cleanup()

	// 準備: Resolverをセットアップ
	resolver := setupTestResolver(t, testDB)

	// 準備: テストデータを挿入
	insertTestLiquorsForResolver(t, resolver.LiquorRepo, 20)

	// テスト実行: 複数回実行
	ctx := context.Background()
	queryResolver := resolver.Query()
	limit := 10

	// 1回目
	result1, err1 := queryResolver.RandomRecommendList(ctx, limit)
	require.Nil(t, err1, "1回目でエラーが発生してはいけません")
	assert.Equal(t, limit, len(result1), "1回目で指定件数が取得できること")

	// 2回目
	result2, err2 := queryResolver.RandomRecommendList(ctx, limit)
	require.Nil(t, err2, "2回目でエラーが発生してはいけません")
	assert.Equal(t, limit, len(result2), "2回目で指定件数が取得できること")

	// 3回目
	result3, err3 := queryResolver.RandomRecommendList(ctx, limit)
	require.Nil(t, err3, "3回目でエラーが発生してはいけません")
	assert.Equal(t, limit, len(result3), "3回目で指定件数が取得できること")

	// すべて正常に動作していることを確認
	t.Logf("複数回の実行がすべて成功しました")
}

// TestRandomRecommendList_GraphQLレスポンスの形式が正しいこと はGraphQLレスポンスの形式をテスト
func TestRandomRecommendList_GraphQLレスポンスの形式が正しいこと(t *testing.T) {
	// 準備: テスト用のデータベースをセットアップ
	testDB, cleanup := setupTestDatabase(t)
	defer cleanup()

	// 準備: Resolverをセットアップ
	resolver := setupTestResolver(t, testDB)

	// 準備: より詳細なテストデータを挿入
	ctx := context.Background()
	collection := resolver.LiquorRepo.DB.Collection(liquorRepository.CollectionName)

	name := "獺祭 純米大吟醸"
	description := "山口県の銘酒。華やかな香りとすっきりとした味わいが特徴です。"
	youtube := "https://www.youtube.com/watch?v=example"
	imageURL := "https://example.com/image.jpg"
	imageBase64 := "base64encodedstring"
	versionNo := 1

	liquor := liquorRepository.Model{
		ID:           primitive.NewObjectID(),
		CategoryID:   1,
		CategoryName: "純米大吟醸",
		Name:         name,
		Description:  &description,
		Youtube:      &youtube,
		ImageURL:     &imageURL,
		ImageBase64:  &imageBase64,
		Rate5Users:   []string{"user1", "user2"},
		Rate4Users:   []string{"user3"},
		Rate3Users:   []string{},
		Rate2Users:   []string{},
		Rate1Users:   []string{},
		UpdatedAt:    time.Now(),
		RandomKey:    0.5,
		VersionNo:    &versionNo,
	}
	_, err := collection.InsertOne(ctx, liquor)
	require.NoError(t, err, "テストデータの挿入に失敗しました")

	// テスト実行: 日本酒リストを取得
	queryResolver := resolver.Query()
	result, err := queryResolver.RandomRecommendList(ctx, 1)

	// 検証: エラーが発生しないこと
	require.Nil(t, err, "エラーが発生してはいけません")
	require.NotNil(t, result, "結果がnilであってはいけません")
	require.Len(t, result, 1, "1件の日本酒が取得できること")

	// 検証: GraphQLレスポンスの各フィールドが正しいこと
	resultLiquor := result[0]

	// 必須フィールド
	assert.Equal(t, liquor.ID.Hex(), resultLiquor.ID, "IDが正しくHex変換されていること")
	assert.Equal(t, liquor.CategoryID, resultLiquor.CategoryID, "カテゴリIDが一致すること")
	assert.Equal(t, liquor.CategoryName, resultLiquor.CategoryName, "カテゴリ名が一致すること")
	assert.Equal(t, liquor.Name, resultLiquor.Name, "名前が一致すること")

	// オプショナルフィールド
	require.NotNil(t, resultLiquor.Description, "説明が設定されていること")
	assert.Equal(t, *liquor.Description, *resultLiquor.Description, "説明が一致すること")

	require.NotNil(t, resultLiquor.Youtube, "YouTube URLが設定されていること")
	assert.Equal(t, *liquor.Youtube, *resultLiquor.Youtube, "YouTube URLが一致すること")

	require.NotNil(t, resultLiquor.ImageURL, "画像URLが設定されていること")
	assert.Equal(t, *liquor.ImageURL, *resultLiquor.ImageURL, "画像URLが一致すること")

	require.NotNil(t, resultLiquor.ImageBase64, "Base64画像が設定されていること")
	assert.Equal(t, *liquor.ImageBase64, *resultLiquor.ImageBase64, "Base64画像が一致すること")

	// 評価ユーザーの配列
	assert.Equal(t, liquor.Rate5Users, resultLiquor.Rate5Users, "評価5のユーザーリストが一致すること")
	assert.Equal(t, liquor.Rate4Users, resultLiquor.Rate4Users, "評価4のユーザーリストが一致すること")
	assert.Equal(t, liquor.Rate3Users, resultLiquor.Rate3Users, "評価3のユーザーリストが一致すること")
	assert.Equal(t, liquor.Rate2Users, resultLiquor.Rate2Users, "評価2のユーザーリストが一致すること")
	assert.Equal(t, liquor.Rate1Users, resultLiquor.Rate1Users, "評価1のユーザーリストが一致すること")

	// 日時フィールド
	assert.NotEmpty(t, resultLiquor.UpdatedAt, "更新日時が設定されていること")

	t.Logf("GraphQLレスポンスのすべてのフィールドが正しく変換されています")
}
