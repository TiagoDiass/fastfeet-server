package repositoryimpl

import (
	"github.com/TiagoDiass/fastfeet-server/internal/entity"
	"gorm.io/gorm"
)

type GormUserRepository struct {
	DB *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) *GormUserRepository {
	return &GormUserRepository{
		DB: db,
	}
}

func (r *GormUserRepository) Create(user *entity.User) error {
	return r.DB.Create(user).Error
}

func (r *GormUserRepository) FindById(id string) (*entity.User, error) {
	var user entity.User

	if err := r.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *GormUserRepository) FindByDocument(document string) (*entity.User, error) {
	var user entity.User

	if err := r.DB.Where("document = ?", document).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *GormUserRepository) Delete(id string) error {
	user, err := r.FindById(id)

	if err != nil {
		return err
	}

	return r.DB.Delete(user).Error
}
