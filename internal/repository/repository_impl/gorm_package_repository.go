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

func (r *GormPackageRepository) FindAllByStatus(status string) ([]*entity.Package, error) {
	var packages []*entity.Package

	err := r.DB.Where("status = ?", status).Order("posted_at asc").Find(&packages).Error

	return packages, err
}

func (r *GormPackageRepository) FindById(id string) (*entity.Package, error) {
	var pkg entity.Package

	err := r.DB.Where("id = ?", id).First(&pkg).Error

	if err != nil {
		return nil, err
	}

	return &pkg, nil
}

func (r *GormPackageRepository) Update(pkg *entity.Package) error {
	return r.DB.Save(pkg).Error
}
