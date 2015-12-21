package author

import (
	"golang.org/x/net/context"

	"github.com/eure/example-blog-golang/model"
)

// Repository is repo for accounts
type Repository struct {
	*model.RootRepository
}

// NewRepository creates new Repository
func NewRepository(ctx context.Context) *Repository {
	return &Repository{
		RootRepository: model.NewRootRepositoryWithSeed(ctx, pkName, new(Entity)),
	}
}

// GetByEmail fetches a single author by email
func (r *Repository) GetByEmail(email string) *Entity {
	ent := new(Entity)
	has, err := r.GetOneByPK(ent, email)
	switch {
	case err != nil:
		return nil
	case !has:
		return nil
	}
	return ent
}
