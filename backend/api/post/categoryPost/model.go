package categoryPost

// RequestData 画像以外の、ShouldBindでバインドするデータ
type RequestData struct {
	Id                *int    `form:"id" binding:"omitempty,gte=1"`
	Name              string  `form:"name" binding:"required,max=100"`
	Parent            int     `form:"parent" binding:"required,gte=1"`
	Description       *string `form:"description" binding:"omitempty,max=5000"`
	VersionNo         *int    `form:"version_no" binding:"omitempty,gte=1"`
	SelectedVersionNo *int    `form:"selected_version_no" binding:"omitempty,gte=1"`
}
