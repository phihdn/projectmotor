package database

import (
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jmoiron/sqlx"
)

type Project struct {
	ID          int32            `db:"id"`
	Title       string           `db:"title"`
	Description pgtype.Text      `db:"description"`
	Published   bool             `db:"published"`
	OwnerID     int32            `db:"owner_id"`
	CreatedAt   pgtype.Timestamp `db:"created_at"`
	UpdatedAt   pgtype.Timestamp `db:"updated_at"`
	Shared      bool             `db:"-"`
}

type ProjectService struct {
	db *sqlx.DB
}

func NewProjectService(db *sqlx.DB) *ProjectService {
	return &ProjectService{
		db: db,
	}
}

func (s *ProjectService) Create(title string, ownerID int32) (Project, error) {
	var project Project
	err := s.db.Get(&project, "INSERT INTO projects (title, owner_id) VALUES ($1, $2) RETURNING *", title, ownerID)
	if err != nil {
		return Project{}, err
	}
	return project, nil
}

func (s *ProjectService) GetAll(ownerID int32) ([]Project, error) {
	var projects []Project
	query := "SELECT * FROM projects WHERE owner_id = $1 ORDER BY created_at desc"
	err := s.db.Select(&projects, query, ownerID)
	if err != nil {
		return []Project{}, err
	}
	return projects, nil
}
