package posttype

import "github.com/divisi-developer-poros/poros-web-backend/config"

type PostTypeInterface interface {
	List() (*[]PostType, error)
	Get(id int) (PostType, error)
	Create(postType *PostType) (*PostType, error)
	Update(postType *PostType) (*PostType, error)
	Delete(id int) error
}

var (
	mysql      config.DBMySQL
	connection = mysql.MysqlConn()
)

func (t *PostType) List() (*[]PostType, error) {
	var postTypes []PostType
	if err := connection.Find(&postTypes).Error; err != nil {
		return nil, err
	}
	return &postTypes, nil
}

func (t *PostType) Get(id uint) (*PostType, error) {
	var postType PostType
	if err := connection.Where("id = ?", id).First(&postType).Error; err != nil {
		return nil, err
	}
	return &postType, nil
}

func (t *PostType) Create(postType *PostType) (*PostType, error) {
	if err := connection.Create(postType).Error; err != nil {
		return nil, err
	}
	return postType, nil
}

func (t *PostType) Update(postType *PostType) (*PostType, error) {
	if _, err := t.Get(postType.ID); err != nil {
		return nil, err
	}

	if err := connection.Save(postType).Error; err != nil {
		return nil, err
	}
	return postType, nil
}

func (t *PostType) Delete(id uint) error {
	if _, err := t.Get(id); err != nil {
		return err
	}
	if err := connection.Delete(&PostType{}, id).Error; err != nil {
		return err
	}
	return nil

}
