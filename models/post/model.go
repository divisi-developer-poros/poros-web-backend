package post

import (
	"github.com/divisi-developer-poros/poros-web-backend/models/postimage"
	"github.com/divisi-developer-poros/poros-web-backend/models/posttype"
	"github.com/divisi-developer-poros/poros-web-backend/models/user"
	"github.com/jinzhu/gorm"
)

// Post ... post model declaration
type Post struct {
	gorm.Model
	Title      string                `json:"title" form:"title" xml:"title" binding:"required"`
	Content    string                `json:"content" form:"content" xml:"content" gorm:"type:longtext" binding:"required"`
	UserID     uint                  `json:"user_id" binding:"-"`
	User       user.User             `json:"user" gorm:"constraint:OnDelete:CASCADE; foreignKey:UserID;" binding:"-"`
	PostTypeID uint                  `json:"post_type_id" form:"post_type_id" xml:"post_type_id" gorm:"column:post_type_id" binding:"required"`
	PostType   posttype.PostType     `json:"post_type" binding:"-"`
	PostImage  []postimage.PostImage `json:"post_image" gorm:"foreignKey:id; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
