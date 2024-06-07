package repository

import "github.com/TiagoDiass/fastfeet-server/internal/entity"

type PackageRepository interface {
	Create(pkg *entity.Package) error
}