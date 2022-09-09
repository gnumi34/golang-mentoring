package books

import (
	"time"

	"gorm.io/gorm"
)

type BookCollections struct {
	BookID      uint           `gorm:"autoIncrement:true, primaryKey, column:book_id" json:"book_id"`
	Title       string         `gorm:"column:title" json:"title"`
	Author      string         `gorm:"column:author" json:"author"`
	Publisher   string         `gorm:"column:publisher" json:"publisher"`
	BookSummary string         `gorm:"column:book_summary" json:"book_summary"`
	Stock       int            `gorm:"column:book_stock" json:"book_stock"`
	MaxStock    int            `gorm:"column:max_book_stock" json:"max_book_stock"`
	Created_At  time.Time      `gorm:"autoCreateTime:true, column:created_at" json:"created_at"`
	Updated_At  time.Time      `gorm:"column:updated_at" json:"updated_at"`
	Deleted_At  gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

type BorrowedBook struct {
	ID         uint              `gorm:"autoIncrement:true, primaryKey, column:id" json:"ID"`
	BookID     uint              `gorm:"foreignKey:BookID, column:book_id" json:"book_id"`
	UserID     uint              `gorm:"foreignKey:UserID, column:user_id" json:"user_id"`
	BorrowedAt time.Time         `gorm:"autoCreateTime:true, column:borrowed_at" json:"borrowed_at"`
	DueDate    time.Time         `gorm:"column:due_date" json:"due_date"`
	IsReturned bool              `gorm:"column:is_returned" json:"is_returned"`
	ReturnedAt time.Time         `gorm:"returned_at" json:"returned_at"`
	Notes      string            `gorm:"column:notes" json:"notes"`
	Books      []BookCollections `gorm:"foreignKey:BookID"`
}

type LendBook struct {
	ID          uint      `gorm:"autoIncrement:true, primaryKey, column:id" json:"ID"`
	BookID      uint      `gorm:"column:book_id" json:"book_id"`
	UserID      uint      `gorm:"column:user_id" json:"user_id"`
	RequestedAt time.Time `gorm:"autoCreateTime:true, column:requested_at" json:"requested_at"`
	IsAccepted  bool      `gorm:"column:is_accepted" json:"is_accepted"`
	Notes       string    `gorm:"column:notes" json:"notes"`
}
