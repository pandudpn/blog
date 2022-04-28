package model

import (
	"database/sql"
	"time"
)

type Blog struct {
	Id        int          `db:"id"`
	CreatedBy int          `db:"created_by"`
	Title     string       `db:"title"`
	Body      string       `db:"body"`
	Image     string       `db:"image"`
	Status    int8         `db:"status"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at"`
}

// IsStatusDraft check status blog is draft
func (b *Blog) IsStatusDraft() bool {
	return b.Status == 0
}

// IsStatusPublish check when status is published
func (b *Blog) IsStatusPublish() bool {
	return b.Status == 1
}

// GetStatus is convert status from db (integer)
// into string for human read able
func (b *Blog) GetStatus() string {
	switch b.Status {
	case 1:
		return "Published"
	default:
		return "Draft"
	}
}
