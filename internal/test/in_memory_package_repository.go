package test

import (
	"github.com/TiagoDiass/fastfeet-server/internal/entity"
)

type InMemoryPackageRepository struct {
	packages []*entity.Package
}

func NewInMemoryPackageRepository() *InMemoryPackageRepository {
	return &InMemoryPackageRepository{
		packages: []*entity.Package{},
	}
}

func (r *InMemoryPackageRepository) FindAllByStatus(status string) ([]*entity.Package, error) {
	availablePackages := []*entity.Package{}
	for _, pkg := range r.packages {
		if pkg.Status == status {
			availablePackages = append(availablePackages, pkg)
		}
	}
	return availablePackages, nil
}

func (r *InMemoryPackageRepository) Create(pkg *entity.Package) error {
	r.packages = append(r.packages, pkg)
	return nil
}
