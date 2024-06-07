package repositoryimpl

import (
	"github.com/TiagoDiass/fastfeet-server/internal/entity"
	"gorm.io/gorm"
)

type GormRecipientRepository struct {
	DB *gorm.DB
}

func NewGormRecipientRepository(db *gorm.DB) *GormRecipientRepository {
	return &GormRecipientRepository{
		DB: db,
	}
}

func (r *GormRecipientRepository) Create(recipient *entity.Recipient) error {
	return r.DB.Create(recipient).Error
}

func (r *GormRecipientRepository) FindById(id string) (*entity.Recipient, error) {
	var recipient entity.Recipient

	if err := r.DB.Where("id = ?", id).First(&recipient).Error; err != nil {
		return nil, err
	}

	return &recipient, nil
}
