package repository

import "github.com/TiagoDiass/fastfeet-server/internal/entity"

type PackageRepository interface {
	Create(pkg *entity.Package) error
	Update(pkg *entity.Package) error
	FindById(id string) (*entity.Package, error)
	FindAllByStatus(status string) ([]*entity.Package, error)
}
