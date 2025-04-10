package db

type User struct {
	ID    uint    `gorm:"primaryKey" json:"id,omitempty" xml:"id" form:"id"`
	Email *string `gorm:"unique;not null" json:"email,omitempty" xml:"email" form:"email"`
}
