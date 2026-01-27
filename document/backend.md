# バックエンド開発ルール
## ディレクトリ構成

特に意識したデザインパターンはないです。開発時にメインで触るのは太字の箇所。

| ディレクトリ | 説明                              |
|------------------|-----------------------------------|
| .devcontainer          | Docker設定が入ってる       |
| api            | RESTで実装したものがこちらに入ってる(主にファイル関連と外部API)                    |
| backend/tmp               | なんかで自動生成されたやつ。デバッガとか？|
| const            | グローバルな定数。共通エラーメッセージとか。    |
| db            | DB関連。リポジトリ層もここで定義。   |
| **db/repository**            | リポジトリ層のロジックはここ。   |
| di            | DI関連。最初よく分からず作ってたので正直中身は要リファクタだと思う。   |
| graph            | GraphQL関連。メインで触るディレクトリの1つ。   |
| graph/generated            |自動生成されたファイルが出力される。触らない。   |
| graph/graphModel            |モデル構造体が自動生成される。触らない。   |
| **graph/resolver**            |MVCでいうCの部分。ここからサービス層を呼び出す。   |
| **graph/schema**            |リゾルバの定義ファイル。   |
|middlewares|ミドルウェア系。認証やエラーログ周り。|
|**service**|一番メインで触る。ビジネスロジック全般を担う。|
|util|その他ユーティリティー系。ヘルパ関数やバリデーターなど。|

なんと**テストは存在しません。** <br>
α版とかがデプロイされたらさすがに作りたい......

## 技術スタック
| 名称 | 説明                              |
|------------------|-----------------------------------|
| Gin          | サーバーフレームワーク       |
| wire          | DIツール       |
| gqlgen          | スキーマファイルからリゾルバを自動生成する       |
|validator| Ginのデフォルト(binding)で対応できないものに対応するために導入されている。基本は使用を避ける。|
|logrus|ログに使用|
|mongo-driver|ORM|

## 基本思想
### リゾルバ(コントローラー)
主にコントローラーとして使い、サービス層を組み合わせて呼び出す。<br>
非常に軽微な処理であれば処理を直書きしても可だが、基本的にはNG。<br>
また、エラーは`*customError.Error`型を返すこと。

### モデル
`repository`配下それぞれの`model.go`に基本的に記載されている。<br>
ドメインとしては同じだが明らかにモデルとして異なる(例えばliquor内の一言掲示板とかはliquorに属するがテーブルとしては別)場合は別のモデルを定義していたりする。(中でディレクトリ切った方がいい説ある...)<br>

### リポジトリ層
モデルと並列に定義されている。DB操作はすべてここに集約する。

### サービス層
`service`配下にわかりやすくまとめる。基本ロジックがここに書かれる。

## 開発手順
### 新規リゾルバ定義時(新しいエンドポイントが欲しい時)
1. `schema/*.graphqls`を編集する。取得系は`Query`、更新系は`Mutation`を`extend`して定義すればOK。<br>
返却型はtypeを使って定義する。<br>
<br>
**⚠️ 重要：スキーマ変更後の必須手順**<br>
スキーマファイル（`schema/*.graphqls`）を変更した場合、フロントエンドの型定義を更新する必要があります。<br>
以下のいずれかの方法で実行してください：<br>
<br>
**方法1：自動スクリプトを使用（推奨）**<br>
```bash
# backend ディレクトリで実行
make update-frontend
```
このコマンドは、バックエンドが起動しているかチェックし、自動的に`fetch-schema`と`codegen`を実行します。<br>
<br>
**方法2：手動で実行**<br>
```bash
# バックエンドを起動（Docker Composeまたは go run main.go）
cd front
npm run fetch-schema  # スキーマを取得（バックエンドが起動している必要があります）
npm run codegen       # 型定義を生成
```
<br>
これを実行しないと、フロントエンドでビルドエラーや型エラーが発生します。<br>
<br>
```graphqls
# typeの例
type ListFromCategory{
  categoryName:String!
  categoryDescription:String
  liquors:[Liquor]!
}

# こんな風に列挙する
extend type Query {
  liquor(id: String!): Liquor!
  randomRecommendList(limit: Int!): [Liquor!]! #ランダムなリスト
  listFromCategory(categoryId: Int!): ListFromCategory! #カテゴリで絞り込んだリスト
  liquorHistories(id: String!):LiquorHistory #編集時に実行する、バージョン履歴つきのデータ
  board(liquorId: String!,page:Int):[BoardPost!]
  getMyBoard(liquorId: String!):BoardPost @optionalAuth #未ログイン時にも呼ばれるのでoptionalに
}

# 更新も同様
extend type Mutation{
  postBoard(input: BoardInput!):Boolean! @optionalAuth
}
```
<br>
2. コードを生成する<br>
`gqlgen generate`で`graph/resolver`配下にリゾルバが自動生成されるので処理を書けばOK。**ルーティングとか意識しなくてOK。**<br>
<br>
3. エラーハンドリング<br>
とりあえず`*customError.Error`型を返せばエラーハンドリングはよしなにやってくれる想定でOK。<br>
基本的には各ロジックファイルに並列で定義されている`errors.go`で管理する。

```go
// こんな風に一意になるエラーコードを定義する
const (
	NotFoundMstData        = "FLAVOR-SERVICE-001-GetFlavorMasterData"
	NotFound               = "FLAVOR-SERVICE-002-NotFound"
	Cursor                 = "FLAVOR-SERVICE-003-Cursor"
	InsertOne              = "FLAVOR-SERVICE-004-InsertOne"
	PostFlavorMapIdFromHex = "FLAVOR-SERVICE-005-PostFlavorMapIdFromHex"
)

func errNotFound(err error, lId primitive.ObjectID, cId int) *customError.Error {
  // 第一引数に生のエラーを渡す。特にエラーオブジェクトがないパターン(論理的エラー)の場合はテキトーに新しく定義してOK。
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusBadRequest, // GraphQLは仕様上全て200で返るため、自分で定義する
		ErrCode:    NotFound, // エラーコードを渡す 
		UserMsg:    errorMsg.DATA, // フロントに表示させたいメッセージを定義。errorMsgは汎用エラーメッセージ。
		Level:      logrus.InfoLevel, // ロガーに対してレベルを定義。InfoLevelならDBに残さない軽微なエラー。ありえない動線のエラーはError以上にしてDBに登録する。
		Input:      fmt.Sprintf("lId: %s, cId: %d", lId.Hex(), cId), // ログに残す目的で入力値を保存しておく。
	})
}
```

基本的にエラーごとに関数を1つ作る構成にしてる。相当ダルいが仕方ない......

## データベース管理

### シーダーデータの更新と反映

シーダーファイル（`db/seeders/categories.json`、`db/seeders/flavorMaps.json`など）を変更した場合、以下の手順でデータベースに反映します。

#### シーダーの実行
```bash
cd backend
make seeder
```

または

```bash
cd backend
go run ./db/seeders/seeder.go
```

シーダーはupsert方式で実装されているため：
- 既存のドキュメント（`id`または`category_id`で判定）がある場合：既存データを更新
- 存在しない場合：新規ドキュメントを作成

#### データベースの完全リフレッシュ

シーダーデータを完全にリセットしたい場合（開発中にデータが壊れた場合など）：

1. MongoDBコンテナを再起動してデータをクリア：
```bash
# ルートディレクトリで実行
docker compose down
docker volume rm sake-db_mongo_data  # ボリュームを削除
docker compose up -d
```

2. シーダーを実行してマスターデータを再投入：
```bash
cd backend
make seeder
```

**注意:** この操作は全てのデータが削除されるため、本番環境では絶対に実行しないでください。

### インデックスの定義と管理

MongoDBのインデックスは`backend/db/indexes/`ディレクトリで管理されています。

#### インデックスの追加手順

1. `backend/db/indexes/define.go`を編集してインデックス定義を追加：

```go
var IndexDefinitions = []IndexDefinition{
    // 既存のインデックス定義...
    
    // 新しいインデックスを追加
    {
        CollectionName: "your_collection_name",
        IndexKeys:      bson.D{{Key: "field_name", Value: 1}}, // 1: 昇順, -1: 降順
        IsNonUnique:    false, // false: ユニーク制約あり, true: 制約なし
        PartialFilter:  bson.D{}, // オプション: 部分インデックスのフィルター条件
    },
}
```

2. バックエンドを再起動：
```bash
go run main.go
```

インデックスは`main.go`の起動時に自動的に作成されます（`indexes.AddIndexes()`関数）。

#### 複合インデックスの例

複数のフィールドを組み合わせたインデックス：

```go
{
    CollectionName: "liquors",
    IndexKeys:      bson.D{{Key: "category_id", Value: 1}, {Key: "created_at", Value: -1}},
    IsNonUnique:    true,
}
```

#### 部分インデックスの例

特定の条件を満たすドキュメントのみにインデックスを適用（null値を除外など）：

```go
{
    CollectionName: "users",
    IndexKeys:      bson.D{{Key: "email", Value: 1}},
    IsNonUnique:    false,
    PartialFilter:  bson.D{{Key: "email", Value: bson.D{{Key: "$exists", Value: true}}}},
}
```

### その他のデータベースTips

#### MongoDB Compassでの確認
- **ダウンロード:** https://www.mongodb.com/try/download/compass
- **接続文字列:** `mongodb://root:root@localhost:27017`

MongoDB Compassを使用すると以下が可能です：
- コレクション内のドキュメントの閲覧・検索
- ドキュメントの追加・編集・削除
- インデックスの確認とパフォーマンス分析
- クエリのパフォーマンステスト

#### パフォーマンスの確認

クエリの実行計画を確認したい場合、MongoDB Compassの「Explain Plan」機能を使用するか、コード内で以下のように実装：

```go
// リポジトリ層でのクエリ実行時
cursor, err := collection.Find(ctx, filter, options.Find().SetHint(bson.D{{Key: "index_name", Value: 1}}))
```

#### トランザクション

複数のコレクションにまたがる操作で整合性を保ちたい場合、MongoDBのトランザクションを使用できます：

```go
session, err := client.StartSession()
if err != nil {
    return err
}
defer session.EndSession(ctx)

err = mongo.WithSession(ctx, session, func(sc mongo.SessionContext) error {
    if err := session.StartTransaction(); err != nil {
        return err
    }
    
    // トランザクション内の処理
    // ...
    
    if err := session.CommitTransaction(sc); err != nil {
        return err
    }
    return nil
})
```
