package test

import (
	"errors"

	"github.com/TiagoDiass/fastfeet-server/internal/entity"
)

var (
	ErrPackageNotExists                 = errors.New("package does not exist")
	ErrOnCreatePackage                  = errors.New("mocked error while creating package")
	NameThatReturnsErrorOnCreatePackage = "Error Package"
)

type InMemoryPackageRepository struct {
	packages []*entity.Package
}

func NewInMemoryPackageRepository() *InMemoryPackageRepository {
	return &InMemoryPackageRepository{
		packages: []*entity.Package{},
	}
}

func (r *InMemoryPackageRepository) Create(pkg *entity.Package) error {
	if pkg.Name == NameThatReturnsErrorOnCreatePackage {
		return ErrOnCreatePackage
	}

	r.packages = append(r.packages, pkg)
	return nil
}

func (r *InMemoryPackageRepository) Update(pkg *entity.Package) error {
	pkgIdx := -1

	for idx, pkgAtIdx := range r.packages {
		if pkgAtIdx.ID == pkg.ID {
			pkgIdx = idx
			break
		}
	}

	if pkgIdx != -1 {
		r.packages[pkgIdx] = pkg

		return nil
	}

	return ErrPackageNotExists
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

func (r *InMemoryPackageRepository) FindById(pkgId string) (*entity.Package, error) {
	for _, pkg := range r.packages {
		if pkg.ID == pkgId {
			return pkg, nil
		}
	}

	return nil, ErrPackageNotExists
}
