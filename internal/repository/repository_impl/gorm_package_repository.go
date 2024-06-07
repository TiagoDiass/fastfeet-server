package repositoryimpl

import (
	"github.com/TiagoDiass/fastfeet-server/internal/entity"
	"gorm.io/gorm"
)

type GormPackageRepository struct {
	DB *gorm.DB
}

func NewGormPackageRepository(db *gorm.DB) *GormPackageRepository {
	return &GormPackageRepository{
		DB: db,
	}
}

func (r *GormPackageRepository) Create(pkg *entity.Package) error {
	return r.DB.Create(pkg).Error
}

func (r *GormPackageRepository) FindAllAvailablePackages() ([]*entity.Package, error) {
	var packages []*entity.Package

	err := r.DB.Order("created_at asc").Find(&packages).Error

	return packages, err
}
