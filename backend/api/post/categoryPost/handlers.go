package categoryPost

import (
	"backend/db/repository/categoriesRepository"
	"backend/db/repository/flavorMapRepository"
	"github.com/aws/aws-sdk-go/service/s3"
	"go.mongodb.org/mongo-driver/mongo"
)

type Handler struct {
	DB                *mongo.Database
	S3Client          *s3.S3
	CategoryRepo      categoriesRepository.CategoryRepository
	FlavorMapMasterRepo flavorMapRepository.FlavorMapMasterRepository
}

// NewHandler 新しいLiquorHandlerを作成するコンストラクタ
func NewHandler(db *mongo.Database, s3Client *s3.S3, categoryRepo categoriesRepository.CategoryRepository, flavorMapMasterRepo flavorMapRepository.FlavorMapMasterRepository) *Handler {
	return &Handler{
		DB:                db,
		S3Client:          s3Client,
		CategoryRepo:      categoryRepo,
		FlavorMapMasterRepo: flavorMapMasterRepo,
	}
}
