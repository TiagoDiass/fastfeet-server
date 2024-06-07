package repository

import "github.com/TiagoDiass/fastfeet-server/internal/entity"

type PackageRepository interface {
	Create(pkg *entity.Package) error
	// FindById(pkgId string) (*entity.Package, error)
	FindAllAvailablePackages() ([]*entity.Package, error)
}
