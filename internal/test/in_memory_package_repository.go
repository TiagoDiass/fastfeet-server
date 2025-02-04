package test

import (
	"errors"

	"github.com/TiagoDiass/fastfeet-server/internal/entity"
)

var (
	ErrPackageNotExists                 = errors.New("package does not exist")
	ErrOnCreatePackage                  = errors.New("mocked error while creating package")
	NameThatReturnsErrorOnCreatePackage = "Package that returns error on create"
	NameThatReturnsErrorOnFindPackages  = "Package that returns error on find"
	ErrOnFindPackages                   = errors.New("mocked error while finding packages")
	NameThatReturnsErrorOnUpdatePackage = "Package that returns error on update"
	ErrOnUpdatePackage                  = errors.New("mocked error while updating package")
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
	if pkg.Name == NameThatReturnsErrorOnUpdatePackage {
		return ErrOnUpdatePackage
	}

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
	packages := []*entity.Package{}
	for _, pkg := range r.packages {
		if pkg.Status == status {
			packages = append(packages, pkg)
		}
	}

	if len(packages) > 0 && packages[0].Name == NameThatReturnsErrorOnFindPackages {
		return nil, ErrOnFindPackages
	}

	return packages, nil
}

func (r *InMemoryPackageRepository) FindById(pkgId string) (*entity.Package, error) {
	for _, pkg := range r.packages {
		if pkg.ID == pkgId {
			return pkg, nil
		}
	}

	return nil, ErrPackageNotExists
}
