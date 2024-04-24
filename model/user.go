package model

type User struct {
	UUIDColumn
	FullName string `gorm:"column:fullName"`
	Email    string `gorm:"column:email"`
	Password string `gorm:"column:password"`
	TimestampColumn
}

func (m *User) TableName() string {
	return "User"
}
