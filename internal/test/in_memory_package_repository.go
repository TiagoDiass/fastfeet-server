package test

import (
	"errors"

	"github.com/TiagoDiass/fastfeet-server/internal/entity"
)

var (
	ErrPackageDoesNotExist = errors.New("package does not exist")
)

type InMemoryPackageRepository struct {
	packages map[string]*entity.Package
}

func NewInMemoryPackageRepository() *InMemoryPackageRepository {
	return &InMemoryPackageRepository{
		packages: make(map[string]*entity.Package),
	}
}

func (r *InMemoryPackageRepository) Create(pkg *entity.Package) error {
	r.packages[pkg.ID] = pkg

	return nil
}

func (r *InMemoryPackageRepository) FindById(pkgId string) (*entity.Package, error) {
	user, exists := r.packages[pkgId]

	if !exists {
		return nil, ErrPackageDoesNotExist
	}

	return user, nil
}
