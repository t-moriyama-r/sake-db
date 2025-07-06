package categoriesRepository

import (
	"backend/graph/graphModel"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	CollectionName     = "categories"
	LogsCollectionName = "categories_logs"
	ID                 = "id"
	Name               = "name"
	ImageURL           = "image_url"
	ImageBase64        = "image_base64"
	Description        = "description"
	Parent             = "parent"
	VersionNo          = "version_no"
	Order              = "order"
	CreateUserId       = "create_user_id"
	CreateUserName     = "create_user_name"
	UpdateUserId       = "update_user_id"
	UpdateUserName     = "update_user_name"
	UpdatedAt          = "updated_at"
	Readonly           = "readonly"
)

// Model 構造体の定義
type Model struct {
	ID             int                 `json:"id" bson:"id"`
	Name           string              `json:"name" bson:"name"`
	Parent         *int                `json:"parent" bson:"parent"`
	Description    *string             `bson:"description"`
	ImageURL       *string             `bson:"image_url"`
	ImageBase64    *string             `bson:"image_base64"`
	VersionNo      *int                `json:"versionNo" bson:"version_no"` //手動で追加したカテゴリはversionNoが存在しない可能性がある
	Children       []*Model            `json:"children,omitempty"`          // 子カテゴリはDBに保存されないため、bsonタグは不要
	Order          *int                `bson:"order"`
	CreateUserId   *primitive.ObjectID `json:"createUserId" bson:"create_user_id"`
	CreateUserName *string             `json:"createUserName" bson:"create_user_name"`
	UpdateUserId   *primitive.ObjectID `json:"updateUserId" bson:"update_user_id"`
	UpdateUserName *string             `json:"updateUserName" bson:"update_user_name"`
	UpdatedAt      time.Time           `json:"updatedAt" bson:"updated_at"`
	Readonly       bool                `bson:"readonly"` //カテゴリ移動不可フラグ
}

func (m *Model) ToGraphQL() *graphModel.Category {
	// 子カテゴリを再帰的に変換
	var children []*graphModel.Category
	if len(m.Children) > 0 {
		children = make([]*graphModel.Category, len(m.Children))
		for i, child := range m.Children {
			children[i] = child.ToGraphQL() // 再帰的にToGraphQLを呼び出す
		}
	}

	var cid, uid *string
	if m.CreateUserId != nil {
		h := (*m.CreateUserId).Hex()
		cid = &h
	}
	if m.UpdateUserId != nil {
		h := (*m.UpdateUserId).Hex()
		uid = &h
	}

	return &graphModel.Category{
		ID:             m.ID,
		Name:           m.Name,
		Parent:         m.Parent,
		Description:    m.Description,
		ImageURL:       m.ImageURL,
		ImageBase64:    m.ImageBase64,
		VersionNo:      m.VersionNo,
		UpdatedAt:      &m.UpdatedAt,
		CreateUserID:   cid,
		CreateUserName: m.CreateUserName,
		UpdateUserID:   uid,
		UpdateUserName: m.UpdateUserName,
		Children:       children, // 変換後の子カテゴリを設定
		Readonly:       m.Readonly,
	}
}
