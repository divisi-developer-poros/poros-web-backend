package posttype

type PostType struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name" form:"name" xml:"name" binding:"required"`
}
