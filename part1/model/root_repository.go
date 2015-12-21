package model

import (
	"fmt"

	"github.com/go-xorm/xorm"
	"golang.org/x/net/context"

	"github.com/eure/example-blog-golang/library/net/context/resource"
)

// RootRepository is root struct for all repositories
type RootRepository struct {
	Ctx    context.Context
	Engine *xorm.Engine

	PrimaryKey string
	Seed       interface{}
	omitList   []string
}

// NewRootRepositoryWithSeed creates new RootRepository
func NewRootRepositoryWithSeed(ctx context.Context, pk string, seed interface{}) *RootRepository {
	return &RootRepository{
		PrimaryKey: pk,
		Ctx:        ctx,
		Engine:     resource.UseDB(ctx),
		Seed:       seed,
	}
}

// ============================================================
//  helper
// ============================================================

// NewSession return session
func (r *RootRepository) NewSession() *xorm.Session {
	return r.Engine.NewSession()
}

// AddOmit adds column name to omit list
func (r *RootRepository) AddOmit(col string) {
	r.omitList = append(r.omitList, col)
}

// GetLastSQL returns the executed SQL statement
func (r *RootRepository) GetLastSQL(s *xorm.Session) string {
	sql, args := s.LastSQL()
	return fmt.Sprintf("%s -- [args] %v", sql, args)
}

// ============================================================
//  create
// ============================================================

// CreateOne inserts new entity data to database
func (r *RootRepository) CreateOne(e interface{}) error {
	s := r.Engine.NewSession()
	if len(r.omitList) > 0 {
		for _, col := range r.omitList {
			s.Omit(col)
		}
		r.omitList = []string{}
	}

	_, err := s.Insert(e)
	if err != nil {
		return err
	}
	return nil
}

// ============================================================
//  get
// ============================================================

// GetOneByPK fetches a single row by primary key
func (r *RootRepository) GetOneByPK(ent, pk interface{}) (bool, error) {
	s := r.Engine.NewSession()
	s.And(r.PrimaryKey+" = ?", pk)
	return r.GetOneBySession(s, ent)
}

// GetOneBySession fetches a single row using the given session
func (r *RootRepository) GetOneBySession(s *xorm.Session, ent interface{}) (bool, error) {
	has, err := s.Get(ent)
	if err != nil {
		return has, err
	}
	return has, nil
}
