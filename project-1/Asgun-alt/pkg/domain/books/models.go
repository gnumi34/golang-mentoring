package books

import (
	"time"

	"gorm.io/gorm"
)

type BookDomain struct {
	BookID      string         `gorm:"primaryKey, column:book_id"`
	Title       string         `gorm:"column:title"`
	Author      string         `gorm:"column:author"`
	Publisher   string         `gorm:"column:publisher"`
	BookSummary string         `gorm:"column:book_summary"`
	Stock       int            `gorm:"column:stock"`
	MaxStock    int            `gorm:"column:max_stock"`
	Created_At  time.Time      `gorm:"column:created_at"`
	Updated_At  time.Time      `gorm:"column:updated_at"`
	Deleted_At  gorm.DeletedAt `gorm:"column:deleted_at"`
}

type BorrowedBook struct {
	ID         string    `gorm:"primaryKey"`
	BookID     string    `gorm:"column:book_id"`
	UserID     string    `gorm:"column:user_id"`
	BorrowedAt time.Time `gorm:"column:borrowed_at"`
	DueDate    time.Time `gorm:"column:due_date"`
	IsReturned bool      `gorm:"column:is_returned"`
	ReturnedAt time.Time `gorm:"returned_at"`
	Notes      string    `gorm:"column:notes"`
}

type LendBook struct {
	ID         string    `gorm:"primaryKey, column:id"`
	BookID     string    `gorm:"column:book_id"`
	UserID     string    `gorm:"column:user_id"`
	RequestAt  time.Time `gorm:"column:request_at"`
	IsAccepted bool      `gorm:"column:is_accepted"`
	Notes      string    `gorm:"column:notes"`
}
